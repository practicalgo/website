set -ex

# Render
hugo

pushd public

# Build
GOOS=linux GOARCH=amd64 go build

# Copy the binary to a temp path 
scp server root@172.105.175.12:/usr/local/bin/practicalgo-website-1

# Atomically replace the existing binary
ssh root@172.105.175.12 mv /usr/local/bin/practicalgo-website-1 /usr/local/bin/practicalgo-website

popd


