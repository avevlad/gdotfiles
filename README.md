# gdotfiles

Automatic detect project language and choice .gitignore from several options (depends on fzf)

## Install

```bash
git clone https://github.com/avevlad/gdotfiles
cd gdotfiles
make build

sudo mv bin/gdotfiles /usr/local/bin
```

## Usage

```bash
Usage:
	gdotfiles [options]

Options:
	--name=...        Name of template (Node | Scala | Symfony | 1C-Bitrix...)
	--type=...        Type of git file (ignore | attributes, default ignore)
	--from=...        Source (github | toptal | local | alexkaratarakis, default
	                  github or alexkaratarakis)
	--verbose         Print all logs output
	--yes             Automatic yes to prompts

Examples:
	# Automatic detect project language and choice .gitignore from several options (depends on fzf)
	gdotfiles

	# Create Scala .gitignore file from github.com/github/gitignore
	gdotfiles --name=Scala --from=github

	# Create C++ .gitattributes file from github.com/alexkaratarakis/gitattributes
	gdotfiles --name=C++ --type=attributes

	# Create two gitignore templates in one .gitignore file from github.com/github/gitignore
	gdotfiles --name=Scala

```