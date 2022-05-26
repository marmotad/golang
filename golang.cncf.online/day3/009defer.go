package main

import "fmt"

//defer关键字用户声明函数，不论函数是否发生错误都在函数执行最后执行(return之前)，若使用defer声明多个函数，则按照声明的顺序，先声明后执行（堆）
//常用来做资源释放，记录日志等工作
func main() {
	defer func() {
		fmt.Println("defer01")
	}()
	defer func() {
		fmt.Println("defer02")
	}()
	fmt.Println("main over")
}
