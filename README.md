# gotest

`github.com/mconbere/gotest` demonstrates an issue with the Go App Engine SDK
and symbolic links. Using a symbolic link in your GOPATH does not work, and
can cause build failures.

A patch is provided which modifies devappserver2 to account for symbolic links
when serving Go code.

# Reproducing the issue

Clone this git repo into a fresh Go Workspace:

    $ mkdir -p work/src
    $ git clone https://github.com/mconbere/gotest work/src/github.com/mconbere/gotest

Fetch go dependencies 
