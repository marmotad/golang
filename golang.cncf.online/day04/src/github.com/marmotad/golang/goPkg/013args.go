package main

import (
	"fmt"
	"os"
)

func main() {
	//获取传递的参数
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
}
