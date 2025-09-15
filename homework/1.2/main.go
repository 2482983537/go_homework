package main

import (
	"fmt"
)

func main() {
	res := isPalindrome(919)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	res := false
	//负数直接返回，结尾0直接返回
	if x < 0 || (x%10 == 0 && x != 0) {
		return res
	}
	revertNum := 0
	// 获取数字长度
	temp := x
	lenX := 0
	for temp > 0 {
		temp /= 10
		lenX++
	}
	y := x
	for i := 0; i < lenX; i++ {
		revertNum = revertNum*10 + y%10
		y /= 10
	}
	if x == revertNum {
		res = true
	}
	return res
}
