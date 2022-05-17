package main

import "fmt"

//定义一个多返回值的函数
func calc1(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

//匿名返回值的函数
func calc2(a, b int) (sum, diff, product, merchant int) {
	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b
	return
}

func main() {
	//使用多个变量接受函数返回值
	a, b, _, _ := calc1(9, 2)
	fmt.Println(a, b)
	fmt.Println(calc2(6, 2))
}
