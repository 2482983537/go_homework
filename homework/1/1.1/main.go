package main

// 初始化 go.mod 文件
// 在项目根目录下运行: go mod init go_homework

import (
	"fmt"
)

func main() {
	res := singleNumber([]int{5, 4, 1, 2, 1, 5, 2, 4, -9, -1, -9})
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	res := 0
	mark := make(map[int]bool)
	for _, val := range nums {
		if mark[val] {
			mark[val] = false
		} else {
			mark[val] = true
		}
	}
	for key, val := range mark {
		if val {
			res = key
		}
	}
	return res
}
