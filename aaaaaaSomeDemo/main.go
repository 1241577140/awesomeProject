package main

import "fmt"

/*
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

示例 1:

输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。
示例 2:

输入: s = "aba"
输出: false
示例 3:

输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)

提示：

1 <= s.length <= 104
s 由小写英文字母组成
*/
func main() {
	fmt.Println(repeatedSubstringPattern("abac"))
}
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		match := false
		for j := i; j < len(s); j++ {
			if s[j] == s[j-i] {
				match = true
			} else {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
