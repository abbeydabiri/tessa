#!/bin/bash
#upx tessa_linux.elf &&
go build -o tessa_linux.elf -ldflags "-s -w" && upx tessa_linux.elf 
