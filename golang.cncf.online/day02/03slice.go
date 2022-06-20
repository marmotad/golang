package main

import "fmt"

func main() {
	//定义切片
	var nums []int
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	//fmt.Printf("%T\n", nums)
	//字面量赋值
	nums = []int{1, 2, 3}
	//fmt.Println(nums)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	//fmt.Println(nums)

	//数组切片赋值
	nums1 := [...]int{1, 2, 3, 4, 5}
	nums = nums1[1:3]
	//fmt.Println(nums)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)

	//make函数声明切片
	//make函数声明切片时，当未指定切片容量时，默认和切片长度一样
	nums = make([]int, 3)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	nums = make([]int, 3, 5)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)

	//切片的操作(CRUD)
	//查切片中内容
	fmt.Println(nums[2])
	//修改切片中的内容
	nums[2] = 100
	fmt.Println(nums[2])
	//增加元素
	nums = append(nums, 4, 5)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	nums = append(nums, 6)
	fmt.Printf("%#v, %d, %d\n", len(nums), cap(nums), nums)
	//遍历数组中的元素
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(i, ":", nums[i])
	}
	for key, value := range nums {
		fmt.Println(key, ":", value)
	}
	//切片的切片操作
	fmt.Printf("%T, %d\n", nums[1:5], nums[1:5])
	//n_cap - start
	n := nums[1:3:10]
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))
	//src_cap - start (原切片长度-新切片的起始长度)
	n = nums[3:4]
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))
	//copy函数，copy(目标切片，原切片)
	nums2, nums3 := []int{1, 2, 4}, []int{10, 20, 30, 40}
	//长度短的切片copy到长度长的切片时，覆盖目标切片前面的值，后面的值不变
	copy(nums3, nums2)
	fmt.Println(nums3)
	//长度长的切片copy到长度短的切片时，覆盖目标切片的值，后面的值忽略
	nums3 = []int{10, 20, 30, 40}
	copy(nums2, nums3)
	fmt.Println(nums2)
	//删除元素
	//删除第一个元素和删除最后一个元素‘
	nums4 := []int{1, 2, 3, 4, 5}
	//删除第一个元素
	fmt.Println(nums4[1:])
	//删除最后一个元素
	fmt.Println(nums4[:len(nums4)-1])
	//删除中间元素（删除nums4的索引为2的元素(值为3)）
	copy(nums4[2:], nums4[3:])
	//copy(nums4[:len(nums4)-1, nums4[0:]])
	fmt.Println(nums4[:len(nums4)-1])
	//堆栈：每次添加在队尾，移除元素在队尾（先进先出）
	//队列：每次添加在队尾，移除元素在队尾（先进后出）
	var queue []int
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	queue = append(queue, 5)
	//1, 2, 3, 4, 5
	fmt.Println(queue)
	//1, 2, 3, 4
	queue = queue[1:]
	fmt.Println(queue)
	//1, 2, 3
	queue = queue[1:]
	fmt.Println(queue)
	//1, 2
	queue = queue[1:]
	fmt.Println(queue)
	//堆栈
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	//1, 2, 3, 4, 5
	fmt.Println(stack)
	//1, 2, 3, 4
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	//1, 2, 3
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	//多维切片
	// 定义多维数组
	var points [][]int
	points1 := make([][]int, 0)
	fmt.Printf("%T\n", points1)
	//对多维数组赋值
	points = append(points, []int{1, 2, 3})
	points = append(points, []int{4, 5, 6})
	points = append(points, []int{7, 8, 9, 10})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0][1])
	//数组是值类型
	slice01 := []int{1, 2, 3}
	slice02 := slice01

	slice02[0] = 10
	fmt.Println(slice01, slice02)
	array01 := [3]int{1, 2, 3}
	array02 := array01
	array02[0] = 10
	fmt.Println(array01, array02)
}
