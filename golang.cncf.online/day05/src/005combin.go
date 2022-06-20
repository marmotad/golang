package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User2 struct {
	ID   int
	Name string
	//定义Addr字段类型为Address结构体
	Addr Address
}

func main() {
	//对被嵌套的结构体赋值
	addr := Address{Region: "beijing", Street: "changping", No: "001"}
	//对User2中的Addr属性赋值
	m1 := User2{Addr: addr}
	fmt.Printf("%#v\n", m1)
	//修改被嵌套的结构体中的值
	m1.ID = 1
	m1.Addr.Region = "china"
	m1.Addr.Street = "changping"
	m1.Addr.No = "9"
	fmt.Println(m1)

}
