#!/bin/bash

# 查找所有包含 go.mod 文件的目录，并在其中存在 Go 源文件时执行 go mod tidy
find . -name "go.mod" -execdir bash -c 'shopt -s nullglob; go_files=(*.go); if [ "${#go_files[@]}" -gt 0 ]; then go mod tidy; fi' \;