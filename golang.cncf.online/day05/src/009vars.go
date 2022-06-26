package main

import "fmt"

type Address4 struct {
	Region string
	street string
	No     string
}

type User6 struct {
	ID   int
	Name string
	Addr Address4
}

type User7 struct {
	ID   int
	Name string
	*Address4
}

//在函数中修改值类型的Name属性的值
func change(u User6) {
	u.Addr.No = "xxxx"
	fmt.Printf("%#v\n", u)
}

//在函数中修改引用类型的Name属性的值
func changePoint(u User7) {
	u.No = "xxxx"
	fmt.Printf("%#v\n", u)
}

func main() {
	me := User6{}
	me2 := me
	me2.Name = "haha"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", me2)
	//定义一个变量，变量类型为User7的结构体
	me3 := User7{
		ID:   1,
		Name: "xixi",
		Address4: &Address4{
			"北京",
			"海淀",
			"100",
		},
	}
	//定义变量me4，并将变量me3赋值给me4
	me4 := me3
	//结构体的值传递
	//修改指针类型的变量值，并打印原变量me3和新变量me4的值
	me4.No = "1000"
	fmt.Println(me3.No)
	fmt.Println(me4.No)

	//在函数体内修改值类型的结构体值，值类型的修改只会在函数内生效
	change(me)
	fmt.Printf("%#v\n", me)
	//在函数体内修指针类型的结构体值，指针类型的修改会在函数外生效
	changePoint(me4)
	fmt.Printf("%#v\n", me4)

}
