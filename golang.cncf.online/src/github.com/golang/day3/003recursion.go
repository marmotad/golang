package main

import "fmt"

//递归函数(自己调用自己)
//使用递归函数计算n个数字之和
func addNum(num int) int {
	// addN(5) => 5 + addN(4)
	// addN(4) => 4 + addN(3)
	// addN(3) => 3 + addN(2)
	// addN(2) => 2 + addN(1)
	// addN(1) => 1
	if num == 1 {
		return 1
	}
	return num + addNum(num-1)
}

//使用递归函数计算n个数字阶乘
func productNum(num int) int {
	if num == 1 {
		return 1
	}
	//定义product保存
	return num * productNum(num-1)
}
func main() {
	fmt.Println(addNum(5))
	fmt.Println(productNum(3))
}
