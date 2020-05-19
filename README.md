# god [![Release](https://img.shields.io/github/release/orhun/God.svg?style=flat-square)](https://github.com/orhun/God/releases)
[![Build Status](https://img.shields.io/travis/orhun/God.svg?style=flat-square)](https://travis-ci.org/orhun/God) [![Go Report Card](https://goreportcard.com/badge/github.com/orhun/god?style=flat-square)](https://goreportcard.com/report/github.com/orhun/god) [![License](https://img.shields.io/badge/license-GPLv3-blue.svg?style=flat-square&color=red)](./LICENSE)

### Linux utility for simplifying the Git usage.

`god` parses the available Git commands from the retrieved list (`git help`) and turns them into an easy-to-type, one or two char format at the execution time.
Shortcuts of [commonly used git commands](https://github.com/joshnh/Git-Commands) are supported for simplifying the usage and speeding up typing even more.

## Installation

### AUR ([god-git](https://aur.archlinux.org/packages/god-git/))

[![AUR](https://img.shields.io/aur/version/god-git.svg?style=flat-square)](https://aur.archlinux.org/packages/god-git/)

Installation can be made with a AUR helper tool like [Trizen](https://aur.archlinux.org/packages/trizen/) or [yay](https://aur.archlinux.org/packages/yay/).

```
trizen god-git
```

### Manual Installation

Install the dependencies.

```
go get -d ./...
```

Set the required Go environment variables if not set for the installation.
See [this link](https://stackoverflow.com/questions/48361893/how-to-set-go-environment-variables-globally) for setting environment variables globally.

```
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

And finally install the `god` package.

```
go install
```

## Commands

```
god > ?
+---------+-----------------------------+
| COMMAND |         DESCRIPTION         |
+---------+-----------------------------+
| exit    | Exit shell                  |
| git     | List available git commands |
| sc      | List git shortcuts          |
| alias   | Print shell or git aliases  |
| help    | Show this help message      |
| version | Show version information    |
| clear   | Clear the terminal          |
+---------+-----------------------------+
```

### **git**

```
god > git
+---------+----------+
| COMMAND |   GIT    |
+---------+----------+
| c       | clone    |
| i       | init     |
| a       | add      |
| m       | mv       |
| r       | reset    |
| rm      | rm       |
| b       | bisect   |
| g       | grep     |
| l       | log      |
| s       | show     |
| st      | status   |
| bn      | branch   |
| ck      | checkout |
| cm      | commit   |
| d       | diff     |
| mr      | merge    |
| ra      | rebase   |
| t       | tag      |
| f       | fetch    |
| p       | pull     |
| ps      | push     |
| mt      | master   |
| o       | origin   |
+---------+----------+
```

_Example output of shortened git commands._

### **sc**

```
god > sc   
+----------+--------------------------------+
| SHORTCUT |            COMMAND             |
+----------+--------------------------------+
| aa       | add -A                         |
| cmt      | commit -m                      |
| rmt      | remote -v                      |
| rr       | rm -r                          |
| ll       | log --graph --decorate --all   |
| lo       | log --graph --decorate         |
|          | --oneline --all                |
| ls       | ls-files                       |
+----------+--------------------------------+
```

_Git shortcuts._

### Executing non-git commands

Other terminal commands can be executed with adding a `'#'` character before the command. It's necessary for non-git commands because the `god` executes all other terminal inputs with `git`.

```
god > # ps
PID   TTY    TIME     CMD
23965 pts/2  00:00:00 bash
30306 pts/2  00:00:00 go
30361 pts/2  00:00:00 god
30519 pts/2  00:00:00 ps
god > # pwd
/home/k3/god
```

## Demo

![demo](https://user-images.githubusercontent.com/24392180/82380773-b2ad5380-9a31-11ea-9579-dab87be71972.gif)

## Example

```sh
god > st
On branch master
	modified:   README.md
no changes added to commit (use "git add" and/or "git commit -a")

god > a README.md

god > cmt "doc: Update README.md"
[master fba7e79] doc: Update README.md
 1 file changed, 14 insertions(+)

god > ps o mt
To https://github.com/orhun/god.git
   45e8aba..fba7e79  master -> master
```

## FAQ

Q: _So it's a tool just to create some shortcuts for Git commands? Have you heard of [git aliases](https://git-scm.com/book/en/v2/Git-Basics-Git-Aliases)?_

A: Yes.

Q: Well, ok.

A: k

## Todo(s)

* Support adding custom shortcuts.
* Auto-complete commands.
* Support custom prompt.

## License

GNU General Public License ([v3](https://www.gnu.org/licenses/gpl.txt))

## Copyright

Copyright (c) 2019-2020, [orhun](https://www.github.com/orhun)