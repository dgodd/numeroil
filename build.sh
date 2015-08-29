#!/usr/bin/env sh

set -e

docker run --rm -v "$PWD":/go/src/github.com/dgodd/numeroil -w /go/src/github.com/dgodd/numeroil golang:1.5 bash -c "go get ./... && go build -o out/numeroil.linux"

ssh wordpress 'mv /home/numeroil/numeroil.linux{,.old}'
scp out/numeroil.linux wordpress:/home/numeroil/numeroil.linux
ssh wordpress 'service numeroil restart'
ssh wordpress 'rm /home/numeroil/numeroil.linux.old'
