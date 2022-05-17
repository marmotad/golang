package main

import "fmt"

//定义一个函数
func typeFunc(a, b int) int {
	return a + b
}

//定义推导长度变量的一个函数
func typeFunc1(a, b int, args ...int) int {
	return a + b
}
func main() {
	//打印函数类型
	fmt.Printf("%T\n", typeFunc)
	fmt.Printf("%T\n", typeFunc1)
	//定义一个变量，将函数类型赋值给变量
	function := typeFunc
	//打印function变量的类型
	fmt.Printf("%T\n", function)
	//调用变量function
	fmt.Println(function(1, 3))
}
