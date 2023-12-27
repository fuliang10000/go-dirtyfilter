package dirtyfilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchFromCustomWords(t *testing.T) {
	tests := []struct {
		replaceChar rune
		input       string
		filterWords []string
		output      string
		keywords    []string
	}{
		{
			replaceChar: '*',
			input:       "我们生活在地球上，每天早上太阳升起，晚上太阳落山，有时候能看见月亮，孩子们看见月亮照在房子上雪白雪白的，脸上露出了笑容。",
			filterWords: []string{"月亮", "太阳", "地球", "房子", "孩子"},
			output:      "我们生活在**上，每天早上**升起，晚上**落山，有时候能看见**，**们看见**照在**上雪白雪白的，脸上露出了笑容。",
			keywords: []string{
				"地球",
				"太阳",
				"太阳",
				"月亮",
				"孩子",
				"房子",
				"月亮",
			},
		},
		{
			replaceChar: '#',
			input:       "We live on Earth, where the sun rises in the morning and sets in the evening. Sometimes we can see the moon, and the children see the moon shining on the house, which is snow-white and snow-white, with a smile on their faces.",
			filterWords: []string{"moon", "sun", "earth", "house", "child", "smile"},
			output:      "We live on #####, where the ### rises in the morning and sets in the evening. Sometimes we can see the ####, and the #####ren see the #### shining on the #####, which is snow-white and snow-white, with a ##### on their faces.",
			keywords: []string{
				"Earth",
				"sun",
				"moon",
				"moon",
				"child",
				"house",
				"smile",
			},
		},
		{
			replaceChar: 'A',
			input:       "我们生活在地球上, where the sun rises in the morning and sets in the evening, 有时候能看见月亮, and the Children see the moon shining on the house, 脸上露出了笑容.",
			filterWords: []string{"月亮", "太阳", "地球", "房子", "孩子", "笑容", "moon", "sun", "earth", "house", "child", "smile"},
			output:      "我们生活在AA上, where the AAA rises in the morning and sets in the evening, 有时候能看见AA, and the AAAAAren see the AAAA shining on the AAAAA, 脸上露出了AA.",
			keywords: []string{
				"地球",
				"sun",
				"月亮",
				"Child",
				"moon",
				"house",
				"笑容",
			},
		},
		{
			replaceChar: '?',
			input:       "我是文明人，不说脏话",
			filterWords: []string{"fuck", "我操"},
			output:      "我是文明人，不说脏话",
			keywords:    []string{},
		},
		{
			replaceChar: '!',
			input:       "测试没有过滤关键字",
			filterWords: []string{},
			output:      "测试没有过滤关键字",
			keywords:    []string{},
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			keywords, output := NewDFAMather().Builder(test.filterWords).Match(test.input, test.replaceChar)
			assert.Equal(t, test.output, output)
			assert.ElementsMatch(t, test.keywords, keywords)
		})
	}
}

func TestMatchFromFileWords(t *testing.T) {
	tests := []struct {
		replaceChar rune
		input       string
		output      string
		keywords    []string
	}{
		{
			replaceChar: '*',
			input:       "毛主席万岁",
			output:      "***万岁",
			keywords: []string{
				"毛主席",
				"主席",
			},
		},
		{
			replaceChar: '#',
			input:       "共产党好啊",
			output:      "###好啊",
			keywords: []string{
				"共产党",
			},
		},
		{
			replaceChar: '.',
			input:       "Fuck you, you are bitch! 我操啊",
			output:      ".... you, you are .....! ..啊",
			keywords: []string{
				"bitch",
				"Fuck",
				"我操",
				"操",
			},
		},
		{
			replaceChar: '*',
			input:       "文明人，不说脏话",
			output:      "文明人，不说脏话",
			keywords:    []string{},
		},
	}
	filterWords := LoadFileWords("./dirty.2017.txt")
	mather := NewDFAMather().Builder(filterWords)
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			keywords, output := mather.Match(test.input, test.replaceChar)
			assert.Equal(t, test.output, output)
			assert.ElementsMatch(t, test.keywords, keywords)
		})
	}
}
