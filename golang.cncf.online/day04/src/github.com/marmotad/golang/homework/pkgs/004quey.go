package pkgs

import (
	"fmt"
	"strings"
)

//查询用户信息
func Quey(users map[string]map[string]string) {
	//定义变量q接受用户查询的内容
	var q string
	fmt.Print("请输入查询信息：")
	fmt.Scan(&q)
	//定义输出格式
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	//打印表头
	fmt.Println(title)
	//使用‘-’格式化
	fmt.Println(strings.Repeat("-", len(title)))
	//遍历map，查找包含变量q的内容
	for _, user := range users {
		fmt.Printf("%5s|%20s|%5s|%20s|%50s\n", user["id"], user["name"], user["age"], user["tel"], user["addr"])
	}
}
