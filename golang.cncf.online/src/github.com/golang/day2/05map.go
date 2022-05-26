package main

import "fmt"

func main() {
	//定义一个映射，默认映射是nil映射
	var scores map[string]int
	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)
	//使用字面量初始化
	scores = map[string]int{"A": 1, "B": 2}
	fmt.Println(scores)
	//使用make函数初始化
	scores = map[string]int{}
	fmt.Println(scores)

	//增删改查
	scores = map[string]int{"A": 1, "B": 2}
	//通过key查询map
	//当查询的key存在时，返回对应的key值，否则返回0
	fmt.Println(scores["A"])
	fmt.Println(scores["C"])
	//通过if判断是否存在
	//方法一
	v, ok := scores["A"]
	if ok {
		fmt.Println(v)
	}
	//方法二
	if v, ok := scores["A"]; ok {
		fmt.Println(v)
	}
	//修改和添加
	//若key不存在，则会添加
	scores["A"] = 10
	scores["C"] = 3
	fmt.Println(scores)
	//删除
	delete(scores, "C")
	fmt.Println(scores)
	//获取映射的长度
	fmt.Println(len(scores))
	//遍历映射
	for k, v := range scores {
		fmt.Println(k, ":", v)
	}
	//映射的映射
	scores01 := map[string]map[string]string{"lisi": {"成绩": "10", "电话": "11", "地址": "中国"}}
	fmt.Println(scores01["lisi"])
	scores01["lisi"] = map[string]string{"成绩": "1000", "电话": "11", "地址": "中国"}
	fmt.Println(scores01)
	scores01["zhangsan"] = map[string]string{}
	fmt.Println(scores01)
	delete(scores01["lisi"], "成绩")
	fmt.Println(scores01)
}
