package main

import (
	"fmt"
)

type Addredd2 struct {
	Region string
	Street string
	No     string
	//ID     int
}

type User5 struct {
	ID   int
	Name string
	//定义Addr字段类型为Address结构体
	Addr Addredd2
}

//匿名定义结构体类型
type Employee3 struct {
	//在结构体中直接引用结构体
	*User5
	Salary float64
	ID     int
}

func main() {
	var me Employee3
	//打印匿名结构体中的指针类型
	fmt.Printf("%#v\n", me)
	//对匿名结构体中的指针类型赋值
	me01 := Employee3{
		User5: &User5{
			1000,
			"xix",
			Addredd2{
				"北京",
				"海淀",
				"111",
			},
		},
		Salary: 1000,
	}

	fmt.Println(me01)
}
