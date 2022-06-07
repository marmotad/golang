package pkgs

import (
	"crypto/md5"
	"fmt"
	"github.com/howeyc/gopass"
)

func AuthUser(PASSWORD string) bool {
	passwordMd5 := md5.Sum([]byte(PASSWORD))
	//var myPassword []byte
	for i := 0; i < 3; i++ {
		fmt.Print("请输入密码：")
		//passwd := gopass.GetPasswd()
		if myPassword, err := gopass.GetPasswd(); err == nil {
			fmt.Println(1, string(myPassword))
			myPasswordMd5 := md5.Sum([]byte(myPassword))
			//fmt.Scan(&myPassword)
			if passwordMd5 == myPasswordMd5 {
				fmt.Println("登录成功！！")
				return true
			}
		}
	}
	defer func() {
		if err := recover(); err != nil {
			return
			fmt.Print(err)
		}
	}()
	return false
}
