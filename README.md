falcon-forwarder
===

## Installation

It is a golang classic project

```bash
# set $GOPATH and $GOROOT
mkdir -p $GOPATH/src/github.com/open-falcon
cd $GOPATH/src/github.com/open-falcon
git clone https://github.com/open-falcon/forwarder.git
cd forwarder
go get ./...
./control build
./control start
```

