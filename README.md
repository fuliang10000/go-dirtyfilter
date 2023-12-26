# go-dirtyfilter

[![Go](https://img.shields.io/badge/Go->=1.21-green)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/fuliang10000/go-dirtyfilter.svg)](https://github.com/fuliang10000/go-dirtyfilter/releases)
[![Report](https://goreportcard.com/badge/github.com/fuliang10000/go-dirtyfilter)](https://goreportcard.com/report/github.com/fuliang10000/go-dirtyfilter)
[![Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/fuliang10000/go-dirtyfilter)
[![License](https://img.shields.io/github/license/fuliang10000/go-dirtyfilter)](https://github.com/fuliang10000/go-dirtyfilter/blob/main/LICENSE)

## 介绍
> 使用go编写的基于Redis实现的分布式锁，lockClient是单例模式，避免资源浪费，锁的操作是线程安全的

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
	filterWords := []string{"朋友", "my friend"}
	text := "你好，我的朋友。Hi my friend, how are you."
	filteredWords, replacedText := dirtyFilter.NewDFAMather().Builder(filterWords).Match(text)
	fmt.Println(filteredWords, replacedText) // [朋友 my friend] 你好，我的**。Hi *********, how are you.
}
```