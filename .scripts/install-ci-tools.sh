#!/bin/bash

set -e

go version
go env -w GO111MODULE=on
go get golang.org/x/lint/golint

cp "$GOPATH"/bin/* /usr/local/bin/
