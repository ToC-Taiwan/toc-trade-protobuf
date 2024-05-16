#!/bin/bash

rm go.mod
rm go.sum

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go mod init github.com/toc-taiwan/toc-trade-protobuf
go mod tidy

git add go.mod go.sum
