#!/bin/bash

cd /Users/guyue/remote/$1
rm -rf .git
git init
git remote add origin $2
git add .
git commit -m "Initial commit"
git push origin master