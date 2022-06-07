package pkgs

import (
	"fmt"
	"strconv"
)

//添加用户
func AddUser(pk int, users map[string]map[string]string) {
	var (
		id   string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Println(id)
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入电话：")
	fmt.Scan(&tel)
	fmt.Print("请输入地址：")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id": id, "name": name, "tel": tel, "age": age, "addr": addr,
	}
	fmt.Println(users)
}
