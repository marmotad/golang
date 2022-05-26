package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 10, 90, 20, 100, -1}
	var maxNum int
	//查找切片中最大的元素
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] > intSlice[i+1] {
			maxNum = intSlice[i]
		}
	}
	//查找切片中第二大的元素
	for i := 0; i < len(intSlice)-1; i++ {
		if intSlice[i] < maxNum {
			if intSlice[i] > intSlice[i+1] {
				fmt.Println(intSlice[i])
			}
		}
	}
}
