package main

import (
	"fmt"
	"time"
)

//定义一个结构体
type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	var me User
	fmt.Println(me)
	fmt.Printf("%#v\n", me)
	//通过字面量对结构体初始化
	var me2 User = User{1, "xix", time.Now(), "china", "111", "xxx"}
	fmt.Printf("%#v\n", me2)
	//使用零值进行初始化
	var me3 = User{}
	fmt.Printf("%#v\n", me3)
	//通过属性名进行初始化,可以省略部分属性值，省略的属性默认值为对应类型的零值
	var me4 = User{ID: 2, Name: "hahaha"}
	fmt.Printf("%#v\n", me4)
	//结构体的指针类型
	fmt.Printf("%f\n%f\n", &me4, &me3)
	//使用结构体的指针类型进行初始化
	var points1 = new(User)
	fmt.Printf("%f\n", points1)
	fmt.Printf("%#v\n", points1)

	var points2 *User = &me3
	fmt.Printf("%f\n", points2)

	var points3 *User = &User{}
	fmt.Printf("%f\n", points3)

	var points4 *User = new(User)
	fmt.Printf("%f\n", points4)
}
