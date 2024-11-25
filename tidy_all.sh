#!/bin/bash

# 查找所有包含 go.mod 文件的目录
find . -name "go.mod" -execdir go mod tidy \;