package main

import (
	"fmt"
	//"golong/golang.cncf.online/pkg/mod/github.com/howeyc/gopass"

	//"golong/golang.cncf.online/day04/pkg/mod/github.com/howeyc/gopass@v0.0.0-20210920133722-c8aef6fb66ef"

	//绝对路径导入
	"golong/golang.cncf.online/day04/src/github.com/marmotad/gopkg"
	//对导入的包定义别名
	localPack "golong/golang.cncf.online/day04/src/github.com/marmotad/gopkg1"
	//。 导入
	. "golong/golang.cncf.online/day04/src/github.com/marmotad/gopkg2"
	//_（空白标识符） 导入，一般与初始化函数一起使用
	_ "golong/golang.cncf.online/day04/src/github.com/marmotad/initpkg"
)

func main() {
	fmt.Println(gopkg.VERSION)
	//调用别名导入
	fmt.Println(localPack.VERSION)
	//调用. 导入
	fmt.Println(VERSION)
	//
	fmt.Println(gopkg.Name)
	fmt.Println("请输入密码：")
	//if bytes, err := gopass.GetPasswd(); err == nil {
	//	fmt.Println(string(bytes))
	//}
}
