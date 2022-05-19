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

//callback 格式化 将传递的数据按照每行打印还是按照一行以|分割打印
func print(callback func(...string), args ...string) {
	fmt.Println("form print function")
	//调用callback函数，可变参数解包后作为callback函数的参数
	callback(args...)
}

func list(args ...string) {
	for k, v := range args {
		fmt.Println(k, ":", v)
	}
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
	//调用print函数，在print函数内部调用list函数对数据进行格式化
	print(list, "A", "B", "C")

	//匿名函数
	sayHello := func(name string) {
		fmt.Println("Hello ", name)
	}
	//调用匿名函数
	sayHello("lisi")
	//定义匿名函数并调用
	func(name string) {
		fmt.Println("Hi", name)
	}("WangWu")
	//使用函数调用新定义的匿名函数
	value := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}
	a := "A"
	print(value, a, "B", "C")

	//函数直接定义匿名函数并调用
	print(func(s ...string) {
		for _, v := range s {
			fmt.Print(v, "\t")
		}
	}, "A", "B", "C")
}
