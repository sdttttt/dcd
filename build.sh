#!/bin/sh

cd cmd

go build -ldflags="-s -w" .
