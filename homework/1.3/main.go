package main

import (
	"fmt"
)

func main() {
	// 测试用例
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"(((",
		")))",
	}
	
	for _, test := range testCases {
		result := isValid(test)
		fmt.Printf("isValid(\"%s\") = %t\n", test, result)
	}
}

func isValid(s string) bool {
	// 创建一个字符栈
	stack := &Stack[rune]{}
	
	// 定义括号映射
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 遇到左括号，入栈
			stack.push(char)
		case ')', '}', ']':
			// 遇到右括号，检查栈顶是否匹配
			if len(stack.items) == 0 {
				return false // 栈为空，没有对应的左括号
			}
			top := stack.pop()
			if top != pairs[char] {
				return false // 括号不匹配
			}
		}
	}
	
	// 最后检查栈是否为空
	return len(stack.items) == 0
}

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) push(item T) {
	stack.items = append(stack.items, item)
}
func (stack *Stack[T]) pop() T {
	var zero T
	if len(stack.items) == 0 {
		return zero
	}
	item := stack.items[len(stack.items)-1]
	index := len(stack.items) - 1
	stack.items = stack.items[:index]
	return item
}
func (stack *Stack[T]) peek() T {
	var zero T
	if len(stack.items) == 0 {
		return zero
	}
	return stack.items[len(stack.items)-1]
}
