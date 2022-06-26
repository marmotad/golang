package main

import (
	"fmt"
	"golong/golang.cncf.online/day05/src/virsibility/user"
)

func main() {
	me := user.User{}
	//访问包里面结构体的属性
	//大写的属性名可以在包外访问，小写的不能在包外访问
	fmt.Println(me.ID)
	//fmt.Println(me.birthDate)

}
