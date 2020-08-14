# go-playpen
Playpen for working with Go

## Setup

### Install Go

```bash
sudo apt-get install golang -y
```

### Environment Variables

```bash
mkdir -p ~/go
```

Ensure ${HOME}/.bash_profile contains GOPATH and PATH contains the bin directory.

```bash
# Source .bashrc
if [ -f "${HOME}/.bashrc" ] ; then
    source "${HOME}/.bashrc"
fi

# Environment Variables
GOPATH-${HOME}/go
PATH=${PATH}:${GOPATH}/bin
```

## Test Your Install

```bash
cd examples/hello
go run hello.go
```

