#!/usr/bin/env bash
protoc -I. --go_out=. --go-grpc_out=. *.proto
mv github.com/reyoung/gt/proto/*.go .
rm -rf github.com
