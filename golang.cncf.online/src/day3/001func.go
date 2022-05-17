package main

import (
	"fmt"
)

//定义一个无参无返回值的函数
func sayHello() {
	fmt.Println("Hello, world!")
}

//定义一个有参数无返回值的函数,name是形参
func sayHi(name string) {
	fmt.Println("Hi," + name)
}

//定义一个有参数有返回值的函数
func add(a int, b int) int {
	return a + b
}

//合并函数参数类型
func add01(a, b int) int {
	return a + b
}

//函数的可变长度参数（可变参数是切片类型） args ...切片类型
func add02(a, b int, args ...int) int {
	sum := a + b
	for _, v := range args {
		sum += v
	}
	return sum
}

//调用其他函数的可变参数
func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		//args... :对args进行解包
		return add02(a, b, args...)
	}
	return -1
}

func main() {
	//打印sayHello的类型
	fmt.Printf("%T\n", sayHello)
	//调用sayHello函数
	sayHello()
	//调用有参数无返回值的函数，lisi是实参
	sayHi("lisi")
	//调用有参数无返回值的函数，lisi是实参
	fmt.Println(add(1, 90))
	fmt.Println(add01(1, 4))
	fmt.Println(add02(1, 4, 5))
	fmt.Println(calc("add", 1, 3, 5))
	fmt.Println(calc("add", 1, 4, 5, 15))
	//通过解包操作删除切片中的元素
	nums := []int{1, 2, 3, 4, 5, 6}
	//删除切片中索引为1的元素
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)
}
