#!/bin/bash

set -e

rm -rf /tmp/git-redlab-tests

mkdir -p /tmp/git-redlab-tests/git-here
mkdir -p /tmp/git-redlab-tests/no-git-here
mkdir -p /tmp/git-redlab-tests/clean-git
mkdir -p /tmp/git-redlab-tests/dirty-git
mkdir -p /tmp/git-redlab-tests/qwerty-branch
mkdir -p /tmp/git-redlab-tests/multiple-branches

cd /tmp/git-redlab-tests/clean-git
git init
echo "test" > abc
git add -A
git commit -m "Test commit"

cd /tmp/git-redlab-tests/dirty-git
git init
echo "test" > abc
git add -A
git commit -m "Test commit"
echo "testtesttest" > abc

cd /tmp/git-redlab-tests/qwerty-branch
git init
git checkout -b qwerty
echo "test" > abc
git add -A
git commit -m "Test commit"

cd /tmp/git-redlab-tests/git-here
git init
git config --add test.abc xyz

cd /tmp/git-redlab-tests/multiple-branches
git init
echo "test" > abc
git add -A
git commit -m "Test commit"
git checkout -b branch1
git checkout -b branch2
git checkout -b branch3
git checkout -b feature/#1

cd /tmp/git-redlab-tests
git clone https://github.com/jereksel/git-redlab-testrepo.git
mv git-redlab-testrepo remote-branches

printf "LINE1" > /tmp/git-redlab-tests/oneline
printf "LINE1\nLINE2" > /tmp/git-redlab-tests/twolines
printf "LINE1\nLINE2\nLINE3" > /tmp/git-redlab-tests/threelines
printf "LINE1\n\nLINE2\n" > /tmp/git-redlab-tests/twolineswithseparation

go test -v github.com/jereksel/git-redlab
