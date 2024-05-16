#!/bin/bash

rm go.mod
rm go.sum

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go mod init gitlab.com/moxa/nbg/nsd/mxview-one-panel-protobuf
go mod tidy

git add go.mod go.sum
