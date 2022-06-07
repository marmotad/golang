package pkgs

import "fmt"

func ModifyUser(id string, users map[string]map[string]string) {
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
