package main

import "fmt"

type User9 struct {
	id   int
	name string
}

//定义函数获取自定义类型的属性
func (user9 User9) GetID() int {
	return user9.id
}

func (user9 User9) GetName() string {
	return user9.name
}

//定义函数修改自定义类型的属性
func (user9 *User9) SetId(id int) {
	user9.id = id
}

func (user9 *User9) SetName(name string) {
	user9.name = name
}

//自定义一个类型，并匿名引用一个其他类型
type Employee4 struct {
	User9
	Salary float64
	name   string
}

func (employee4 Employee4) GetName() string {
	return employee4.name
}

func (employee4 *Employee4) SetName(name string) {
	employee4.name = name
}
func main() {
	var me = Employee4{
		User9: User9{
			1,
			"kkkkk",
		},
		Salary: 1000,
		name:   "laozhang",
	}
	//获取自定义属性中的值
	//修改相应属性值时，会优先寻找最近的属性的方法，如果想要使用父属性的方法需要全路径调用
	fmt.Println(me.User9.GetName())
	fmt.Println(me.GetName())
	//修改自定义属性中的值
	me.SetName("haha")
	fmt.Println(me.GetName())

}
