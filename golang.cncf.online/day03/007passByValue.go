package main

import "fmt"

//定义一个参数为int类型的函数
func changeInt(a int) {
	//在函数内修改a的值
	a = 100
	fmt.Println("in function :", a)
}

//定义一个参数为int切片类型的函数
func changeSlice(s []int) {
	//在函数内修改s[0]的值
	s[0] = 100
	fmt.Println("in function :", s)
}

//指针类型的值传递
func changeIntByPoint(p *int) {
	*p = 1000
	fmt.Println("in function :", p)

}

func main() {
	//值类型不会修改原变量的值
	num := 10
	changeInt(num)
	fmt.Println(num)

	//引用类型会修改原变量的值
	nums := []int{1, 2, 3, 4, 5}
	changeSlice(nums)
	fmt.Println(nums)
	//指针类型会修改原变量的值
	//打印原指针地址
	fmt.Println(&num)
	changeIntByPoint(&num)
	fmt.Println(num)
	fmt.Println(&num)
}
