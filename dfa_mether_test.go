package DFAMather

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMatchChinese 测试中文过滤
func TestMatchChinese(t *testing.T) {
	filterChineseWords := []string{"月亮", "太阳", "地球", "房子", "孩子"}
	chineseText := "我们生活在地球上，每天早上太阳升起，晚上太阳落山，有时候能看见月亮，孩子们看见月亮照在房子上雪白雪白的，脸上露出了笑容。"
	filteredWords, replacedText := NewDFAMather().Builder(filterChineseWords).Match(chineseText)
	assert.Equal(t, len(filteredWords), 7)
	assert.Equal(t, replacedText, "我们生活在**上，每天早上**升起，晚上**落山，有时候能看见**，**们看见**照在**上雪白雪白的，脸上露出了笑容。")
}

// TestMatchEnglish 测试英文过滤
func TestMatchEnglish(t *testing.T) {
	filterEnglishWords := []string{"moon", "sun", "earth", "house", "child", "smile"}
	englishText := "We live on Earth, where the sun rises in the morning and sets in the evening. Sometimes we can see the moon, and the children see the moon shining on the house, which is snow-white and snow-white, with a smile on their faces."
	filteredWords, replacedText := NewDFAMather().Builder(filterEnglishWords).Match(englishText)
	assert.Equal(t, len(filteredWords), 7)
	assert.Equal(t, replacedText, "We live on *****, where the *** rises in the morning and sets in the evening. Sometimes we can see the ****, and the *****ren see the **** shining on the *****, which is snow-white and snow-white, with a ***** on their faces.")
}

// TestMatchMixLanguage 测试混合语言过滤
func TestMatchMixLanguage(t *testing.T) {
	filterWords := []string{"月亮", "太阳", "地球", "房子", "孩子", "笑容", "moon", "sun", "earth", "house", "child", "smile"}
	text := "我们生活在地球上, where the sun rises in the morning and sets in the evening, 有时候能看见月亮, and the children see the moon shining on the house, 脸上露出了笑容."
	filteredWords, replacedText := NewDFAMather().Builder(filterWords).Match(text)
	assert.Equal(t, len(filteredWords), 7)
	assert.Equal(t, replacedText, "我们生活在**上, where the *** rises in the morning and sets in the evening, 有时候能看见**, and the *****ren see the **** shining on the *****, 脸上露出了**.")
}
