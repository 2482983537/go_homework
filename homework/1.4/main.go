package main

import (
	"fmt"
)

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1))
}

func longestCommonPrefix(strs []string) string {
	//判空
	if len(strs) == 0 {
		return ""
	}
	//判断第一个字符
	head := ""

	for i := 0; i < len(strs[0]); i++ {
		char := []rune(strs[0])[i]
		for j := 0; j < len(strs); j++ {
			//超长也退出
			if i > len(strs[j]) {
				return head
			}
			//有不相等退出
			if char != []rune(strs[j])[i] {
				return head
			}
		}
		head += string(char)
	}
	return head
}
