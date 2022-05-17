package main

import "fmt"

func main() {
	//定义多维数组
	marrays := [3][2]int{}
	fmt.Println(marrays)
	//访问多维数组中的元素
	fmt.Println(marrays[0])
	//修改多维数组中的数组值
	marrays[0] = [2]int{1, 2}
	fmt.Println(marrays[0])
	//修改多维数组中数组的值
	marrays[0][1] = 1000
	fmt.Println(marrays[0])
	//多维数组的字面量
	marrays = [3][2]int{{11, 12}, {21, 22}}
	fmt.Println(marrays)
}
