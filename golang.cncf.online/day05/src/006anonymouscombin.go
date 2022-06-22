package main

import "fmt"

type Address1 struct {
	Region string
	Street string
	No     string
	//ID     int
}

type User3 struct {
	ID   int
	Name string
	//定义Addr字段类型为Address结构体
	Addr Address1
}

//匿名定义结构体类型
type Employee2 struct {
	//在结构体中直接引用结构体
	User3
	Salary float64
	ID     int
}

func main() {

	me := Employee2{}
	fmt.Printf("%T\n%#v\n", me, me)

	me2 := Employee2{
		User3: User3{
			1,
			"haha",
			Address1{
				"北京",
				"海淀",
				"上地",
			},
		},
		Salary: 1000,
	}
	fmt.Printf("%T\n%#v\n", me2, me2)
	//访问结构中的值
	fmt.Println(me2.User3.Name)
	//直接访问匿名嵌入的机构体中的参数
	fmt.Println(me2.Name)
	//当结构体中的名称与被引用的匿名结构体中的名称相同时，优先访问结构体中的
	fmt.Println(me2.ID)
	fmt.Println(me2.User3.ID)
	me2.ID = 10
	fmt.Println(me2.ID)
	//当属性名相同时，必须使用全路径访问
	//fmt.Printf("%T\n%#v\n", me2, me2)

}
