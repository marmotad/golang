package main

import "fmt"

func main() {
	//要排序的数组
	array := []int{10, 100, 90, 88, 46}
	//var sortArray []int{}
	//临时变量存贮最大值
	var tmpNum int
	//遍历数组中的每个值
	for i := 0; i < len(array)-1; i++ {
		//如果当前值最大
		if array[i] >= tmpNum {
			tmpNum = array[i]
		}
	}
	fmt.Println(tmpNum)
}
