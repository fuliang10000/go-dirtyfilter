package dirtyfilter

import (
	"bufio"
	"os"
	"strings"
)

// LoadFileWords 读取一个换行符分割的文本文件
func LoadFileWords(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	var filterWords []string
	// 逐行扫描
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words := strings.ToLower(strings.TrimSpace(scanner.Text())) // 均处理为小写
		if words != "" {
			filterWords = append(filterWords, words)
		}
	}

	return filterWords
}
