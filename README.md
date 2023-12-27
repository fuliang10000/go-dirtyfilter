# go-dirtyfilter

[![Go](https://img.shields.io/badge/Go->=1.21-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/fuliang10000/go-dirtyfilter.svg)](https://github.com/fuliang10000/go-dirtyfilter/releases)
[![Report](https://goreportcard.com/badge/github.com/fuliang10000/go-dirtyfilter)](https://goreportcard.com/report/github.com/fuliang10000/go-dirtyfilter)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/fuliang10000/go-dirtyfilter)
[![License](https://img.shields.io/github/license/fuliang10000/go-dirtyfilter)](https://github.com/fuliang10000/go-dirtyfilter/blob/main/LICENSE)

## 介绍
> 使用go语言编写的基于DFA算法实现的敏感词过滤器，为降低trie树的体积，起到同时过滤大小写的效果，过滤单词请统一使用小写，
> 附带脏词基础库：[dirty.2017.txt](dirty.2017.txt)

## 快速开始

### 安装
```bash
go get -u github.com/fuliang10000/go-dirtyfilter
```

### 测试
```bash
go test github.com/fuliang10000/go-dirtyfilter  -v -cover
```

### Use Demo
```go
package main

import (
	"fmt"
	dirtyFilter "github.com/fuliang10000/go-dirtyfilter"
)

func main() {
	filterWords := dirtyFilter.LoadFileWords("./dirty.2017.txt")
	text := "毛主席万岁！！"
	replaceChar := '*'
	filteredWords, replacedText := dirtyFilter.NewDFAMather().Builder(filterWords).Match(text, replaceChar)
	fmt.Println(filteredWords, replacedText) // [毛主席 主席] ***万岁！！
}
```