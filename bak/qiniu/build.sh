#!/bin/sh

GOOS=linux GOARCH=amd64 go build upload.go
scp upload liuwei@moon:/home/liuwei
