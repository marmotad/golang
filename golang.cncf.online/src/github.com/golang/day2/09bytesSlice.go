package main

import (
	"bytes"
	"fmt"
)

func main() {
	var bytes01 []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	fmt.Printf("%T, %#v\n", bytes01, bytes01)
	//转换成字符串类型
	s := string(bytes01)
	fmt.Printf("%T, %#v\n", s, s)
	//字符串转换成byte
	bs := []byte(s)
	fmt.Printf("%T, %#v\n", bs, bs)

	//字节切片的函数
	fmt.Println(bytes.Compare([]byte("abc"), []byte("xyz")))
	fmt.Println(bytes.Compare([]byte("abc"), []byte("def")))
	fmt.Println(bytes.Index([]byte("abcdefabc"), []byte("def")))
}
