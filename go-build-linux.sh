#!/bin/bash
#upx tessa_linux.elf &&
GOOS=linux GOARCH=amd64 go build -o tessa_linux.elf -ldflags "-s -w" #&& upx tessa_linux.elf 
