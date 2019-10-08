#!/bin/bash
#go build  -o tessa_mac.app -ldflags "-s -w" && mv tessa_mac.app app/.
#go build  -o tessa_mac.app -ldflags "-s -w" && upx "-9" tessa_mac.app && mv tessa_mac.app app/.
go build  -o tessa_mac.app -ldflags "-s -w"
