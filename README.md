# gh-repos
A small tool to get the names of repos from a Github's user or org.

Pre-requisites:

* You need to config an env called GH_AUTH_TOKEN with your personal access token, to do the requests

Get the name of all public repos from **edivangalindo**

```bash
echo edivangalindo | gh-repos
```

Get name of all public repos from a list of users

```bash
cat users.txt | gh-repos
```

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile gh-repos:
```
go install github.com/edivangalindo/gh-repos@latest
```

You can now run `~/go/bin/gh-repos`. If you'd like to just run `gh-repos` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.
