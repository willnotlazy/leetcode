package questions

import "math"

/**
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
*/
func MinWindows(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)

	start, left, right, valid, length := 0, 0, 0, 0, math.MaxInt

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for right < len(s) {
		c := s[right]
		right ++
		if _, ok := need[c]; ok {
			window[c] ++
			if window[c] == need[c] {
				valid ++
			}
		}
        
		for valid == len(need) {
			d := s[left]
			if length > right - left {
				start = left
				length = right - left
			}
			left ++
			if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
				window[d]--
			}
		}
	}

	if length > len(s) {
		return ""
	}

	return s[start: start + length]
}