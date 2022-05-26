package main

import "fmt"

func main() {
	var nums [10]int
	//定义一个长度为10的数组
	fmt.Printf("%T \n", nums)
	//打印变量类型
	fmt.Println(nums)
	//数组的元素的零值就是对应数据类型的零值
	nums = [10]int{1, 2, 3}
	//对数组中元素赋值，默认从索引为0的元素依次赋值
	fmt.Println(nums)
	nums = [10]int{9: 9}
	//通过索引对指定元素赋值
	fmt.Println(nums)
	nums1 := [...]int{1, 2, 3}
	//定义一个数组，数组长度根据数组中元素的值进行推倒
	fmt.Printf("%T \n", nums1)
	nums2 := [...]int{1, 2, 3, 5: 10}
	//定义一个数组，数组长度根据数组中元素的值进行推倒
	fmt.Printf("%T \n", nums2)

	//操作
	nums3 := [...]int{1, 3, 4}
	nums4 := [...]int{1, 3, 5}
	//判断数组是否相等
	fmt.Println(nums3 == nums4)
	//获取数组长度,数组的长度一旦定义之后就不能改变
	fmt.Println(len(nums3))
	fmt.Println(nums3[(len(nums3) - 1)])

	//遍历数组
	for i := 0; i < len(nums3); i++ {
		fmt.Println(i, ":", nums3[i])
	}

	for index, value := range nums3 {
		fmt.Println(index, ":", value)
	}
	//切片
	fmt.Println(nums3)
	//定义数组的切片
	fmt.Printf("%T\n", nums3[0:2])
	//数组切片的容量最大只能和数组长度相同，若切片的容量小于切片的元素数会舍弃末尾的元素
	fmt.Println(nums3[0:2:2])
}
