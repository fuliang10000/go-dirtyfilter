package DFAMather

import (
	"sync"
	"unicode"
)

// Node 代表DFA的一个节点。
type Node struct {
	End  bool           // 是否为一个单词的结束。
	Next map[rune]*Node // 一个映射，用于存储此节点的所有子节点。
}

// DFAMatcher 代表一个完整的DFA。
type DFAMatcher struct {
	root *Node      // DFA的根节点。
	mu   sync.Mutex // 互斥锁，保证线程安全
}

// NewDFAMather 创建出一个DFA树的根节点实例
func NewDFAMather() *DFAMatcher {
	return &DFAMatcher{
		root: &Node{
			End: false,
		},
	}
}

// Builder 用于构建DFA，将提供的所有单词添加到DFA中。
func (d *DFAMatcher) Builder(words []string) *DFAMatcher {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, item := range words {
		d.root.addWord(item)
	}
	return d
}

// addWord 方法用于向当前节点添加一个单词。
// 这个方法会遍历单词的每一个字符，并为每一个字符添加一个子节点。
func (n *Node) addWord(word string) {
	node := n // 从当前节点开始。
	chars := []rune(word)
	for index := range chars {
		node = node.addChild(chars[index])
	}
	node.End = true // 设置最后一个节点为单词的结束。
}

// addChild 向当前节点添加一个子节点。
// 如果子节点已经存在，它将不会被重复添加。
func (n *Node) addChild(c rune) *Node {
	// 如果Next字段为nil，则初始化一个映射。
	if n.Next == nil {
		n.Next = make(map[rune]*Node)
	}
	// 检查字符c是否已经是当前节点的子节点。
	// 如果ok为true，则字符c已经是当前节点的子节点，直接返回该子节点。
	if next, ok := n.Next[c]; ok {
		return next
	} else {
		// 否则，创建一个新的节点，并将其设置为当前节点的子节点。
		n.Next[c] = &Node{
			End:  false,
			Next: nil,
		}
		return n.Next[c] // 返回新创建的子节点。
	}
}

// Match 在文本中查找并替换敏感词。
// 它返回找到的敏感词列表和替换后的文本。
func (d *DFAMatcher) Match(text string, replaceChar rune) (sensitiveWords []string, replaceText string) {
	// 如果DFA是空的，说明敏感词库为空，直接返回原始文本。
	if d.root == nil {
		return nil, text
	}
	textChars := []rune(text)
	textCharsCopy := make([]rune, len(textChars))
	copy(textCharsCopy, textChars)
	length := len(textChars)
	for i := 0; i < length; i++ {
		// 在DFA树中查找当前字符对应的子节点
		temp := d.root.findChild(textChars[i])
		if temp == nil {
			continue
		}
		j := i + 1
		// 遍历文本中的字符，查找匹配的敏感词，第一个匹配上了，就进行后面的向下匹配
		for ; j < length && temp != nil; j++ {
			if temp.End {
				// 如果找到一个敏感词，将其添加到结果列表中，并在副本中替换为指定字符
				sensitiveWords = append(sensitiveWords, string(textChars[i:j]))
				d.replaceRune(textCharsCopy, replaceChar, i, j) //替换敏感词
			}
			temp = temp.findChild(textChars[j])
		}
		// 处理文本末尾的情况，如果末尾是一个完整的敏感词，添加到结果列表中，并在副本中替换为指定字符
		if j == length && temp != nil && temp.End {
			sensitiveWords = append(sensitiveWords, string(textChars[i:length]))
			d.replaceRune(textCharsCopy, replaceChar, i, length)
		}
	}
	return sensitiveWords, string(textCharsCopy) // 返回匹配到的敏感词列表和替换后的文本
}

// findChild 用于在当前节点的子节点中查找一个特定的子节点。
func (n *Node) findChild(c rune) *Node {
	// 如果Next字段为nil，则直接返回nil。
	if n.Next == nil {
		return nil
	}
	// 将大写字母转换为小写，便于同时匹配大小写并过滤
	if unicode.IsUpper(c) {
		c = unicode.ToLower(c)
	}
	// 检查字符c是否是当前节点的子节点。
	if _, ok := n.Next[c]; ok { // 如果ok为true，则字符c是当前节点的子节点，返回该子节点。
		return n.Next[c]
	}
	return nil // 否则，返回nil。
}

// replaceRune 替换掉文本切片中出现的关键词
func (d *DFAMatcher) replaceRune(chars []rune, replaceChar rune, begin int, end int) {
	for i := begin; i < end; i++ {
		chars[i] = replaceChar
	}
}
