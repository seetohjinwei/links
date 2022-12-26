#!/bin/sh

cd /root/git/links

git pull

# build Go app
go build -o bin/links .
echo 'Built Go app!'

# restart service
service links restart
