package main

import "fmt"

type Addr struct {
	Region string
	Street string
	No     string
}
type User4 struct {
	ID   int
	Name string
	Addr *Addr
}

func main() {
	var me01 User4
	fmt.Printf("%#v\n", me01)

	me02 := User4{
		ID:   1,
		Name: "haha",
		//使用指针类型的结构体
		Addr: &Addr{"北京", "昌平区", "001"},
	}
	fmt.Println(me02)
	//访问指针类型结构体的数据
	fmt.Println(me02.Addr.Region)
	//修改指针类型结构体的数据
	me02.Addr.No = "10"
	fmt.Println(me02.Addr.No)

}
