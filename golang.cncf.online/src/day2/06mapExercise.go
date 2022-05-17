package main

import "fmt"

func main() {
	users := []string{"A", "C", "B", "C", "D", "E", "A", "A"}
	votes := map[string]int{}
	for _, v := range users {
		//遍历数组，获取数组中的value（数组中每个索引的值）
		votes[v] += 1
		//对映射votes添加元素（map中不存在的值为0，添加直接将值改为1，如果存在则对值+1）
	}
	fmt.Println(votes)
}
