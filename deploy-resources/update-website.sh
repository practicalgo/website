#!/bin/bash

set -ex

# Render
hugo

pushd public

# Build
GOOS=linux GOARCH=amd64 go build

# Copy the binary to a temp path 
scp server root@practicalgobook.net:/usr/local/bin/practicalgo-website-1

ssh root@practicalgobook.net systemctl stop practicalgo-website
ssh root@practicalgobook.net mv /usr/local/bin/practicalgo-website-1 /usr/local/bin/practicalgo-website
ssh root@practicalgobook.net systemctl start practicalgo-website
popd


