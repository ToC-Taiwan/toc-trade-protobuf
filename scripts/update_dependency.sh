#!/bin/bash
set -e

rm -rf node_modules
rm -rf package-lock.json

npm i
npm i -g protoc-gen-ts
npm i -g npm-check-updates
ncu --peer -u

npm install

git add package.json
git add package-lock.json
