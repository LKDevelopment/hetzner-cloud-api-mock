#!/bin/bash -e
gofmt -l $(find . -type f -name '*.go' | grep -v ./vendor/)
