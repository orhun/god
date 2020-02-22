package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/carmark/pseudo-terminal-go/terminal"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"os/exec"
	"strings"
)

var version = "1.7"
var outb, errb bytes.Buffer                 // Buffer for command output
var cmdSlice, gitCmdSlice, allCmds []string // Slice for the shortened git commands
var gitShortcuts [][]string                 // 2-d slice for the git shortcuts
var whiteColor (*color.Color) = color.New(color.FgWhite, color.Bold)
var restartTerm bool = false // Handling the stdout issues.
var termChar = "#"           // Character for non-git terminal commands.
var promptStr = "[god ~]$ "

// Executes the terminal command and returns output.
// stdout parameter determines the output stream.
func execCmd(input string, stdout bool) string {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")
	// Prepare the command to execute.
	// sh used for handling the command parameters.
	// Otherwise, exec library takes the parameters
	// as argument which is something that we don't
	// want due to the complexity of git commands.
	cmd := exec.Command("sh", "-c", input)
	// Set the correct output device.
	if stdout {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	} else {
		cmd.Stdout = &outb
		cmd.Stderr = &errb
	}
	// Execute the command and return output
	// depending on the stdout parameter.
	cmd.Run()
	if !stdout {
		return outb.String()
	}
	return ""
}

// Search the given query in slice.
// Returns true if the element exists.
func searchInSlice(slice []string, query string) bool {
	set := make(map[string]bool)
	for _, v := range slice {
		set[v] = true
	}
	return set[query]
}

// Returns a slice with given dimension parameter.
// Used for getting keys or values from a 2-d slice.
func getShortcutSlice(slice [][]string, d int) []string {
	var shortcuts []string
	for _, shortcut := range slice {
		shortcuts = append(shortcuts, shortcut[d])
	}
	return shortcuts
}

// Prepare (shorten) the git commands.
func prepareCmds() []string {
	// Show status if repository exists in directory.
	execCmd("git status", true)
	// Trimming the string using sed.
	// 's/^\s*//' -> Substitute the found expression (' ' with '').
	// 's/ *[A-Z].*//' -> Remove the git command description
	// && -> For the next command.
	removeSpaces := "sed -e 's/^\\s*//' -e 's/ *[A-Z].*//' && "
	// Parsing the git commands.
	// grep '^  *[a-z]' -> Select the lines starting with indent.
	// tr -d '*' -> Remove the '*' character.
	parseGitCmd :=
		"git help | grep '^  *[a-z]' | " + removeSpaces +
			"git branch | tr -d '*' | " +
			strings.Replace(removeSpaces, " -e 's/ *[A-Z].*//'", "", 1) +
			"git remote"
	cmdStr := execCmd(parseGitCmd, false)
	gitCmdSlice = strings.Split(cmdStr, "\n")
	for i, cmd := range gitCmdSlice {
		if len(cmd) > 0 {
			// Use the first character of git command
			// for the new command if not exists in the
			// commands slice. (cmdSlice)
			// If first character is in the list, compose
			// a two character abbreviation for it and
			// add it to slice.
			firstChar := string([]rune(cmd)[0])
			if !searchInSlice(cmdSlice, firstChar) {
				cmdSlice = append(cmdSlice, firstChar)
			} else {
				cmdSlice = append(cmdSlice, firstChar+
					string([]rune(cmd)[len(cmd)/2]))
			}
		} else {
			// Remove empty character
			gitCmdSlice = append(gitCmdSlice[:i], gitCmdSlice[i+1:]...)
		}
	}
	// Add git shortcuts.
	gitShortcuts = append(gitShortcuts,
		[]string{"add -A", "aa"},
		[]string{"commit -m", "cmt"},
		[]string{"remote -v", "rmt"},
		[]string{"rm -r", "rr"},
		[]string{"log --graph --decorate --all", "ll"},
		[]string{"log --graph --decorate --oneline --all", "lo"},
		[]string{"ls-files", "ls"},
		[]string{"config credential.helper store", "crds"})
	// Create a slice for storing all commands.
	allCmds = append(gitCmdSlice, getShortcutSlice(gitShortcuts, 0)...)
	return gitCmdSlice
}

// Create a git command from the given string.
// Returns changed/new command.
func buildCmd(line string) string {
	line = " " + line + " "
	// Run command without git if it starts
	// with '#' character.
	if strings.Contains(string([]rune(line)[:2]), termChar) {
		return strings.Replace(line, " "+termChar, " ", 1)
	}
	// Support the commands starting with git.
	line = strings.Replace(line, " git ", " ", -1)
	// Use the first quoted part if command contains quote.
	quotedPart := ""
	if strings.Contains(line, "\"") {
		quotedPart = strings.Join(strings.Split(line, "\"")[1:], "\"")
		line = strings.Split(line, "\"")[0] + "\""
	}
	// Replace the shortened command with its original.
	for index, cmd := range append(cmdSlice, getShortcutSlice(gitShortcuts, 1)...) {
		cmd = " " + cmd + " "
		if strings.Contains(line, cmd) {
			line = strings.Replace(line, cmd, " "+allCmds[index]+" ", -1)
		} else if strings.Contains(line, strings.ToUpper(cmd)) {
			line = strings.Replace(line, strings.ToUpper(cmd), " "+allCmds[index]+" ", -1)
		}
	}
	// Add quoted part back to the command.
	if quotedPart != "" {
		line += quotedPart
	}
	return "git" + line
}

// Start the interactive shell.
func startTerm(persistent bool) {
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut()
	whiteColor.Println("Type '?' for help or 'git' for list of commands.")
	term.SetPrompt(promptStr)
cmdLoop:
	for {
		// Read the keyboard input.
		line, err := term.ReadLine()
		// Exit on Ctrl-D and Ctrl-C (if -p not provided).
		if (err == io.EOF) || (line == "^C" && !persistent) {
			fmt.Println()
			return
		}
		// Built-in commands.
		switch line {
		case "", " ":
			break
		case "clear":
			execCmd("clear", true)
		case "exit":
			break cmdLoop
		case "?", "help":
			showHelp()
		case "version":
			showVersion()
		case "git":
			showCommands()
		case "sc":
			showShortcuts()
		case "alias":
			promptAlias()
		default:
			// Build the git command.
			gitCmd := buildCmd(line)
			// Release the std in/out for preventing the
			// git username & password input issues.
			if strings.Contains(gitCmd, "push") {
				restartTerm = true
				term.ReleaseFromStdInOut()
			}
			// Handle the execution of the input.
			if retval := execCmd(gitCmd, true); len(retval) > 0 {
				fmt.Fprintln(os.Stderr, retval)
			}
			// Restart the terminal for flushing the stdout.
			// It is necessary for input required situations.
			if restartTerm {
				term, _ = terminal.NewWithStdInOut()
				defer term.ReleaseFromStdInOut()
				term.SetPrompt(promptStr)
				restartTerm = false
			}
		}
	}
}

// Takes 'table' parameter and returns colored.
func setTableColors(table *tablewriter.Table) *tablewriter.Table {
	whiteTable := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgHiWhiteColor}
	blackTable := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgHiBlackColor}
	table.SetHeaderColor(whiteTable, whiteTable)
	table.SetColumnColor(whiteTable, blackTable)
	return table
}

// Display help message.
func showHelp() {
	cliCmds := map[string]string{
		"git":     "List available git commands",
		"sc":      "List git shortcuts",
		"alias":   "Print shell or git aliases",
		"help":    "Show this help message",
		"version": "Show version information",
		"clear":   "Clear the terminal",
		"exit":    "Exit shell"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "Description"})
	table = setTableColors(table)
	for k, v := range cliCmds {
		table.Append([]string{k, v})
	}
	table.Render()
}

// Show git commands in table.
func showCommands() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "git"})
	table = setTableColors(table)
	for index, cmd := range cmdSlice {
		table.Append([]string{cmd, gitCmdSlice[index]})
	}
	table.Render()
}

// Show commonly used git commands shortcuts.
func showShortcuts() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Shortcut", "Command"})
	table = setTableColors(table)
	for _, shortcut := range gitShortcuts {
		table.Append([]string{shortcut[1], shortcut[0]})
	}
	table.Render()
}

// Save shortened commands as shell alias or Git alias.
func promptAlias() {
	aliasOpts := []string{
		"shell",
		"git",
	}
	formats := map[string]string{
		"shell": "alias %s='git %s'",
		"git":   "%s = %s",
	}
	aliasPrompt := promptui.Select{
		Label: "Create alias for",
		Items: aliasOpts,
	}
	_, selection, err := aliasPrompt.Run()
	if err != nil {
		fmt.Printf("Selection failed: %v\n", err)
	}
	if len(selection) > 0 {
		whiteColor.Println("Alias list for " + selection + ":")
		for index, cmd := range append(cmdSlice, getShortcutSlice(gitShortcuts, 1)...) {
			alias := fmt.Sprintf(formats[selection], cmd, allCmds[index])
			fmt.Println(alias)
		}
	}
}

// Show project information including version.
func showVersion() {
	fmt.Println()
	asciiFigure := figure.NewFigure("god", "cosmic", true)
	asciiFigure.Print()
	whiteColor.Println("\n ~ god:v" + version)
	fmt.Println(" ~ utility for simplifying the git usage" +
		"\n ~ github.com/orhun/god" +
		"\n ~ licensed under gplv3\n")
}

func main() {
	// If -v argument is given, show project information and exit.
	// If not, start the god terminal.
	versionFlag := flag.Bool("v", false, "Show version information")
	persistentFlag := flag.Bool("p", false, "Don't exit on ^C")
	flag.Parse()
	if *versionFlag {
		showVersion()
	} else {
		prepareCmds()
		startTerm(*persistentFlag)
	}
}
