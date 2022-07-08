package questions

import (
	"bytes"
)

// ReverseWords
// 557. 反转字符串中的单词 III
// 给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//示例 1：
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//
//示例 2:
//
//输入： s = "God Ding"
//输出："doG gniD"
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/reverse-words-in-a-string-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func ReverseWords(s string) string {
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		if sByte[i] == ' ' {
			continue
		}
		j := i + 1
		for ; j < len(sByte); j++ {
			if sByte[j] == ' ' {
				break
			}
		}
		tmp := j + 1
		j = j - 1
		for ; i < j; i++ {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			j--
		}
		i = tmp
	}

	return string(sByte)
}

func ReverseWords2(s string) string {
	i := 0
	j := 0
	n := len(s)
	ret := make([]byte, n)
	for ; j < n; j++ {
		if s[j] == ' ' {
			for k := j - 1; k >= i; k-- {
				ret = append(ret, s[k])
			}
			ret = append(ret, ' ')
			i = j + 1
		}
	}
	if i != j {
		for k := j - 1; k >= i; k-- {
			ret = append(ret, s[k])
		}
	}
	return string(ret)
}

func replaceSpace(s string) string {

	var count int
	for i := 0; i > len(s); i++ {
		if s[i] == ' ' {
			count++
		}
	}

	buffer := bytes.NewBuffer(make([]byte, 0, len(s)+count*2))
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			buffer.WriteByte(s[i])
		} else {
			buffer.Write([]byte{'%', '2', '0'})
		}
	}

	return buffer.String()
}
