package main

import (
	"fmt"
	"sort"
)

func main() {
	//对数字类型切片排序
	intSlice := []int{1, 6, 10, 20}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	//对字符串类型的切片排序
	strSlice := []string{"A", "c", "ef", "ddr"}
	sort.Strings(strSlice)
	fmt.Println(strSlice)
	//对小数类型切片排序
	floatSlice := []float64{1, 1.2, 1.12}
	sort.Float64s(floatSlice)
	fmt.Println(floatSlice)
}
