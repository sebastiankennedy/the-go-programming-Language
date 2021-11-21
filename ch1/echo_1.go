package main

import (
	"fmt"
	"os"
)

func main() {
	// 声明字符串变量，隐式赋予其类型的 Zero Value
	var s, sep string

	// os.Args 变量是一个字符串的切片，切片是一个简化版的动态数组
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}