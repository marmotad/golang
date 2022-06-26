package main

import "fmt"

//定义一个结构体
type Dog struct {
	Name string
}

//为结构体Dog定义方法call
func (dog Dog) Call() {
	fmt.Printf("%s:宝宝在干嘛吃早饭了吗 没吃的话我给你买\n", dog.Name)
}

//为结构体的属性值无法通过变量赋值方式修改
func (dog Dog) SetName(name string) {
	dog.Name = name
}

//通过指针类型对结构体的属性值进行修改
func (dog *Dog) PointSetName(name string) {
	dog.Name = name
}

func main() {
	dog := Dog{"老李"}
	dog.Call()
	dog.SetName("老张")
	dog.Call()

	//定义一个变量，类型为结构体的指针类型
	pdog := &Dog{"老赵"}
	//通过取引用的方法调用Call函数
	(*pdog).Call()
	//通过变量的地址修 改结构体属性的值
	//解引用
	(*pdog).PointSetName("老刘")
	//自动解引用（语法糖）
	(*pdog).Call()

	//取引用
	(&dog).PointSetName("小李")
	//自动取引用（语法糖）
	dog.PointSetName("小刘")
	dog.Call()

	//不能直接调用指针类型的方法
	var pdog2 *Dog
	fmt.Println(pdog2.Call)
}
