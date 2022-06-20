package main

import "fmt"

//匿名定义结构体类型
type Employee2 struct {
	//在结构体中直接引用结构体
	User2
	Salary float64
}

func main() {

	m2 := Employee2{}
	fmt.Printf("%T\n%#v\n", m2, m2)
}
