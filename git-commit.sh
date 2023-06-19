#!/bin/bash
set -o errexit

git pull

git add .
git commit -m "."
git push