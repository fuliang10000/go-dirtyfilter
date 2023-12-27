package dirtyfilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLoadFileWords 测试读取文本文件
func TestLoadFileWords(t *testing.T) {
	path := "./dirty.2017.txt"
	words := LoadFileWords(path)
	assert.Equal(t, 5767, len(words))
}
