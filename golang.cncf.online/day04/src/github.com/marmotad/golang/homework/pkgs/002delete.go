package pkgs

import "fmt"

func DeleteUser(id string, users map[string]map[string]string) {
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
