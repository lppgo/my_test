#!/bin/bash

echo -e "\033[0;32mDeploying updates to <gitlab.yuliangtec.cn/luna/com.yuliangtec.luna.proto>\033[0m"

# Build the project.

echo -e "\033[0;32mgo mod tidy\033[0m"
go mod tidy

echo -e "\033[0;32mgit status\033[0m"
git status

echo -e "\033[0;32mgit stash\033[0m"
git stash

echo -e "\033[0;32mgit pull\033[0m"
git pull

echo -e "\033[0;32mgit stash pop\033[0m"
git stash pop

echo -e "\033[0;32mgit status\033[0m"
git status

echo -e "\033[0;32mgit add -A\033[0m"
git add -A

# # Commit changes.
msg="rebuilding site $(date)"
if [ $# -eq 1 ]; then
    msg="$1"
fi
echo -e "\033[0;32mgit commit -m "$msg"\033[0m"
git commit -m "$msg"

# Push source and build repos.
echo -e "\033[0;32mgit push origin master\033[0m"
git push origin master
