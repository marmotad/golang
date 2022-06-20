package main

import "fmt"

func main() {
	//声明匿名结构体
	var me struct {
		ID   int
		Name string
	}
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)
	//简短声明一个结构体
	me2 := struct {
		ID   int
		Name string
	}{1, "xixi"}
	fmt.Printf("%#v\n", me2)
}
