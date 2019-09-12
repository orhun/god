# god [![Release](https://img.shields.io/github/release/orhun/God.svg?style=flat-square)](https://github.com/orhun/God/releases)
[![AUR](https://img.shields.io/aur/version/god-git.svg?style=flat-square)](https://aur.archlinux.org/packages/god-git/) [![Build Status](https://img.shields.io/travis/orhun/God.svg?style=flat-square)](https://travis-ci.org/orhun/God) [![Go Report Card](https://goreportcard.com/badge/github.com/orhun/god?style=flat-square)](https://goreportcard.com/report/github.com/orhun/god) [![License](https://img.shields.io/badge/license-GPLv3-blue.svg?style=flat-square&color=red)](./LICENSE)

### Linux utility for simplifying the Git usage.

`god` parses the available Git commands from the retrieved list (`git help`) and turns them into an easy-to-type, one or two char format at the execution time.
Shortcuts of [commonly used git commands](https://github.com/joshnh/Git-Commands) are supported for simplifying the usage and speeding up typing even more.

## Installation

### AUR ([god-git](https://aur.archlinux.org/packages/god-git/))

Installation can be made with a AUR helper tool like [Trizen](https://aur.archlinux.org/packages/trizen/) or [yay](https://aur.archlinux.org/packages/yay/).

```
trizen god-git
```

![Trizen Installation](https://user-images.githubusercontent.com/24392180/58751001-95d6e380-84a1-11e9-9314-888c94ab0475.gif)

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
[god ~]$ ?
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
[god ~]$ git
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
[god ~]$ sc   
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
[god ~]$ # ps
PID   TTY    TIME     CMD
23965 pts/2  00:00:00 bash
30306 pts/2  00:00:00 go
30361 pts/2  00:00:00 god
30519 pts/2  00:00:00 ps
[god ~]$ # pwd
/home/k3/god
```

## Demo

![v1.0-demo](https://user-images.githubusercontent.com/24392180/58592279-c97ef700-8270-11e9-8290-862ab278ca4b.gif)

## Example

```
[god ~]$ st
On branch master
	modified:   README.md
no changes added to commit (use "git add" and/or "git commit -a")

[god ~]$ a README.md

[god ~]$ cmt "README.md updated~"
[master fba7e79] README.md updated~
 1 file changed, 14 insertions(+)

[god ~]$ ps o mt
To https://github.com/orhun/God.git
   45e8aba..fba7e79  master -> master
```

## FAQ

Q: _So it's an application just to create some shortcuts to git commands? Did you heard of [git aliases](https://git-scm.com/book/en/v2/Git-Basics-Git-Aliases)?_

A: The project aims to provide a interactive shell that have built-in shortened Git commands and shorcuts. Benefits of it depends on your workspace, people like to use Git in command line and work with multiple shells can use it as a helper tool. Unlike the git aliases, it only needs one execution and there will be no need for typing 'git' anymore for using git. Here's a screenshot of my workspace while using `god` in a separate terminal.

![Workspace](https://user-images.githubusercontent.com/24392180/58703963-166ae680-83b3-11e9-8e3b-1b92b70a7583.jpg)

## Todo(s)

* Support adding custom shortcuts.
* Terminal autocomplete feature.

## License

GNU General Public License v3. (see [gpl](https://www.gnu.org/licenses/gpl.txt))

## Credit

Copyright (C) 2019 by orhun https://www.github.com/orhun
