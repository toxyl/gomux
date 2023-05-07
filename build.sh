#!/bin/bash
go mod tidy && CGO_ENABLED=0 go build -o gomux . && mv gomux /usr/local/bin/ && gomux

