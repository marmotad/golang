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
			modifyUser(string(id), users)
		case "3":
			deleteUser(string(id), users)
		case "4":
			quey(users)
		//如果输入5退出函数
		case "5":
			break
		}
	}
}

func deleteUser(id string, users map[string]map[string]string) {
	fmt.Print("请输入要删除的用户ID：")
	fmt.Scan(&id)
	var yes string
	fmt.Print("是否确定修改用户信息（Y|N）：")
	fmt.Scan(&yes)
	if yes == "y" || yes == "Y" || yes == "yes" || yes == "YES" {
		delete(users, id)
	} else if yes != "y" {
		fmt.Println("用户取消修改，修改未生效")
	}
	fmt.Println(users)
}

func modifyUser(id string, users map[string]map[string]string) {
	fmt.Print("请输入要修改的用户ID：")
	fmt.Scan(&id)
	//定义一个变量用于接收需要修改的类型
	var modifyKey string
	fmt.Print("请输入要修改的用户属性：")
	fmt.Scan(&modifyKey)
	//输入新的值。并使用一个变量进行存储
	fmt.Print("请输入新的值：")
	var newValue string
	fmt.Scan(&newValue)
	//根据modifyKey的值判断修改的类型
	var yes string
	fmt.Print("是否确定修改用户信息（Y|N）：")
	fmt.Scan(&yes)
	if yes == "y" || yes == "Y" || yes == "yes" || yes == "YES" {
		if modifyKey == "name" {
			users[id]["name"] = newValue
		} else if modifyKey == "age" {
			users[id]["age"] = newValue
		} else if modifyKey == "tel" {
			users[id]["tel"] = newValue
		} else if modifyKey == "addr" {
			users[id]["addr"] = newValue
		} else {
			fmt.Println("输入的修改内容错误，请重新输入！！！")
		}
	} else if yes != "y" {
		fmt.Println("用户取消修改，修改未生效")
	}
	fmt.Println(users)
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
