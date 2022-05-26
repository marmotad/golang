package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//声明一个map存储用户信息
	users := make(map[string]map[string]string)
	id := 0
	fmt.Print("欢迎使用用户管理系统！！！")
	for {
		// 定义一个string 类型变量op
		var op string
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
请输入指令:`)
		//从标准输入输入字符串
		fmt.Scan(&op)
		switch op {
		case "1":
			addUser(id, users)
			id++
		case "2":

		case "3":

		case "4":
			quey(users)
		//如果输入5退出函数
		case "5":
			break
		}
	}
}

//查询用户信息
func quey(users map[string]map[string]string) {
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

//添加用户
func addUser(pk int, users map[string]map[string]string) {
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
