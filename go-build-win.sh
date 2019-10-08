#!/bin/bash
GOOS=windows GOARCH=386 go build  -o tessa_win.exe -ldflags "-s -w" && upx tessa_win.exe
