package main

import "fmt"

func main() {
	//定义一个返回值是函数类型的函数
	addBase := func(base int) func(int) int {
		//返回一个函数类型
		return func(i int) int {
			return base + i
		}
	}
	//定义一个函数调用addBase函数
	add10 := addBase(10)
	fmt.Println(add10(1))
	//直接调用addBase函数
	fmt.Println(addBase(5)(1))
}
