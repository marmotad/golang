package main

import (
	"fmt"
	"homework04/pkgs"
)

func main() {
	if pkgs.AuthUser("123456") == true {
		id := 0
		users := make(map[string]map[string]string)
		fmt.Print("欢迎使用用户管理系统！！！")
		for {
			// 定义一个string 类型变量op
			var op string
			fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 打印所有用户信息
6. 退出
请输入指令:`)
			fmt.Scan(&op)
			switch op {
			case "1":
				pkgs.AddUser(id, users)
				id++
			case "2":
				pkgs.ModifyUser(string(id), users)
			case "3":
				pkgs.DeleteUser(string(id), users)
			case "4":
				pkgs.Quey(users)
			case "5":
				pkgs.PrintAll(users)
			//如果输入5退出函数
			case "6":
				break
			}
		}
	}
}
