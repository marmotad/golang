package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//计算byte类型md5，并使用大写格式输出
	fmt.Printf("%X\n", md5.Sum([]byte("i am lisi")))
	bytes := md5.Sum([]byte("i am zhangsan"))
	//将byte类型的md5s值转换为string[] 类型
	fmt.Println(hex.EncodeToString(bytes[:]))
	//计算大文件md5值
	MD5Sum := md5.New()
	MD5Sum.Write([]byte("i am"))
	MD5Sum.Write([]byte(" zhangsan"))
	fmt.Printf("%x\n", MD5Sum.Sum(nil))
}
