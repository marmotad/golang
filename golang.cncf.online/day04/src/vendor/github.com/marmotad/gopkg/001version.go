package gopkg

import "fmt"

const VERSION float64 = 6.11

//Go 语言使用名称首字母大小写来判断对象( 常量、变量、函数、类型、结构体、方法等)的访问权限，首字母大写标识包外可见 (公开的)，否者仅包内可访问(内部的)
const Name, name = "haha", "xixi"

func Test() {
	fmt.Println("Test")
}
