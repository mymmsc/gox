#!/usr/bin/env bash

set -e

git remote set-url origin https://gitee.com/quant1x/gox.git
git checkout master
git fetch --all
git pull --all
git remote -vv