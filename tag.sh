#!/bin/bash 

git tag
git tag v"$*" 
git push origin --tags 
git tag
git show v"$*"

