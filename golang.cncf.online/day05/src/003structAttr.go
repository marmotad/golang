package main

import (
	"fmt"
	"time"
)

//定义一个结构体
type User1 struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	var userName User1
	//通过属性名访问
	fmt.Println(userName.ID)
	//通过属性名修改值
	userName.ID = 1
	fmt.Println(userName.ID)
	//通过指针类型修改和访问数据类型
	me2 := &User1{
		ID:   20,
		Name: "haha",
	}
	fmt.Println(me2)
	//通过指针类型访问
	(*me2).Tel = "11122223333"
	//通过变量名访问
	me2.Addr = "china"
	fmt.Println(me2)
}
