#!/usr/bin/env sh

go generate ./... > /dev/null 2>&1
go build
