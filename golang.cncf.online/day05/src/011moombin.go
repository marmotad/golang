package main

import "fmt"

type Address5 struct {
	Region string
	Street string
	No     string
}

//定义一个函数，以指定输出格式
func (addr Address5) String() string {
	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.No)
}

type User8 struct {
	ID   int
	Name string
	Addr Address5
}

func (u User8) String() string {
	//fmt.Sprintf 格式化字符串
	return fmt.Sprintf("[%d]%s: %s", u.ID, u.Name, u.Addr)
}

func main() {
	var u User8 = User8{ID: 1, Name: "John", Addr: Address5{"北京", "昌平", "111"}}
	fmt.Println(u)
	fmt.Println(u.Addr)
}
