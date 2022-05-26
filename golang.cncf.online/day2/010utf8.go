package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱中国"
	//获取utf8字符数量
	fmt.Println(utf8.RuneCountInString(s))
}
