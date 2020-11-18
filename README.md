# Network-latency

Network-latency is a project created for demonstrating the difference in actual data bytes using different protocols.
It uses a sample data comprising name and email of a person to demonstrate the difference of data packet size using HTTP 1.1, HTTP 2 and gRPC.

## Setup

This service runs on `go`.

- Install go
    - On OSX run `brew install go`.
    - Follow instructions on `https://golang.org/doc/install` for other OSes.
- Setup go
    - Make sure that the executable `go` is in your shell's path.
    - Add the following in your .zshrc or .bashrc: (where `< workspace_dir >` is the directory in which you'll checkout your code)
```
GOPATH= < workspace_dir >
export GOPATH
PATH="${PATH}:${GOPATH}/bin"
export PATH
```

- Checkout the code under GOPATH directory, install the dependencies and build the project:
```
$ git clone git@github.com:kamalashreen/network-latency.git
$ cd <path-to-network-latency>
```

## Usage
- Build the project
```
$ go build -o network-latency
```
- Run the binary by passing corresponding protocol name to get the byte size for given data.
1. For HTTP
```
$ ./network-latency http
```
2. For HTTP2
```
$ ./network-latency http2
```
3. For gRPC
```
$ ./network-latency grpc
```
