# Passphrase

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/milligan22963/passphrase)](https://github.com/milligan22963/passphrase/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/milligan22963/passphrase.svg)](https://pkg.go.dev/github.com/milligan22963/passphrase)
[![go.mod](https://img.shields.io/github/go-mod/go-version/milligan22963/passphrase)](go.mod)
[![LICENSE](https://img.shields.io/github/license/milligan22963/passphrase)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/milligan22963/passphrase/build)](https://github.com/milligan22963/passphrase/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/asphaltbuffet/passphrase)](https://goreportcard.com/report/github.com/asphaltbuffet/passphrase)
[![Codecov](https://codecov.io/gh/asphaltbuffet/passphrase/branch/main/graph/badge.svg)](https://codecov.io/gh/asphaltbuffet/passphrase)

## Description

Generate a random passphrase inspired by the [XKCD comic](https://xkcd.com/936/). The default word list is an edited subset of data from [wordfrequency.info](https://www.wordfrequency.info/).

## Usage

## Setup

Below you can find sample instructions on how to set up the development environment.
Of course, you can use other tools like [GoLand](https://www.jetbrains.com/go/),
[Vim](https://github.com/fatih/vim-go), [Emacs](https://github.com/dominikh/go-mode.el).
However, take notice that the Visual Studio Go extension is
[officially supported](https://blog.golang.org/vscode-go) by the Go team.

1. Install [Go](https://golang.org/doc/install).
1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
1. Clone and open this repository.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version.
More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher),
[Go Blog](https://blog.golang.org/v2-go-modules).

## Note on building in Windows

### How can I build on Windows

Install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/)
and copy `C:\TDM-GCC-64\bin\mingw32-make.exe`
to `C:\TDM-GCC-64\bin\make.exe`.
Alternatively, you may install [mingw-w64](http://mingw-w64.org/doku.php)
and copy `mingw32-make.exe` accordingly.

Take a look [here](https://github.com/docker-archive/toolbox/issues/673#issuecomment-355275054),
if you have problems using Docker in Git Bash.

You can also use [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
or develop inside a [Remote Container](https://code.visualstudio.com/docs/remote/containers).
However, take into consideration that then you are not going to use "bare-metal" Windows.

Consider using [goyek](https://github.com/goyek/goyek)
for creating cross-platform build pipelines in Go.

## Contributing

Feel free to create an issue or submit a PR.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
