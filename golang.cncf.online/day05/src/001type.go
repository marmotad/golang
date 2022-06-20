package main

import "fmt"

//自定义一个基本数据类型
type Counter int

//自定义一个符合数据类型
type user map[string]string

//自定义一个函数数据类型
type callback func(...string)

func main() {
	//定义一个类型
	var counter Counter
	counter += 10
	fmt.Println(counter)
	counter = 130
	fmt.Println(counter)

	myUser := make(user)
	myUser["name"] = "xixi"
	myUser["age"] = "19"
	fmt.Println(myUser)
	fmt.Printf("%T, %T\n", counter, myUser)

	var testPrint callback = func(users ...string) {
		fmt.Println(users)
	}
	testPrint("test func type")
	//自定义数据类型的类型转换和比较
	int1 := 178
	var int2 Counter = 111
	fmt.Println(int1 > int(int2))
	fmt.Println(Counter(int1) > int2)
}
