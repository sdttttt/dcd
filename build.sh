#!/bin/sh

cd cmd

go build -ldflags="-s -w" .

mv ./cmd/huc* .