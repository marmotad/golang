package main

import "fmt"

//a:源，b：借助。c:目的
func tower(a, b, c string, later int) {
	if later == 1 {
		fmt.Println(a, "->", c)
		return
	}
	//a n -1 借助 c 移动到 b
	tower(a, c, b, later-1)
	//b n -1 借助 a 移动到 c
	tower(b, a, c, later-1)
}

func main() {
	//三层的
	tower("A", "B", "C", 3)
}
