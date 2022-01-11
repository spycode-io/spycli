#!/bin/bash
version=$1

env GOOS=linux GOARCH=amd64 go build . && zip -q -r spycli-linux-amd64-$version.zip spycli README.md assets/docs assets/img && rm spycli
env GOOS=darwin GOARCH=amd64 go build . && zip -q -r spycli-darwin-amd64-$version.zip spycli README.md assets/docs assets/img && rm spycli
env GOOS=windows GOARCH=amd64 go build . && zip -q -r spycli-windows-amd64-$version.zip spycli.exe README.md assets/docs assets/img && rm spycli.exe
