#!/bin/bash

# This shell is executed before docker build.

dockerName=blog_server
version=v1.0
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

sudo docker build --platform linux/amd64 -t $dockerName:$version .
