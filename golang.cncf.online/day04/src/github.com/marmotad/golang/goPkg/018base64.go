package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//使用base64编码
	base := base64.StdEncoding.EncodeToString([]byte("i am lisi"))
	fmt.Println(base)
	//对base64编码的内容进行解码
	unbase, err := base64.StdEncoding.DecodeString(base)
	//将解码后的byte数组转换为string类型并打印
	fmt.Println(string(unbase), err)
}
