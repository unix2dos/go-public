#!/bin/sh


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
scp sub_header root@tencent2:/root



