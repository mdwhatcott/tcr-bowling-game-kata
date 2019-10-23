#!/usr/bin/env bash

set -e

TITLE='tcr-bowling-game-kata'

echo '=> Resetting...'
rm -rf *.go go.mod go.sum .idea/ .git/

echo '=> Initializing Go module...'
go mod init github.com/mdwhatcott/$TITLE
go get -u github.com/smartystreets/gunit

echo '=> Initializing git repository...'
git init
git add .
git commit -m "Initial commit."
git remote add origin git@github.com:mdwhatcott/$TITLE.git

echo '=> Pushing initial commit...'
git push -f origin master

echo '=> Starting IDE...'
goland . --wait

echo '=> Pushing final state...'
git push -f origin master

echo '=> Finished.'
