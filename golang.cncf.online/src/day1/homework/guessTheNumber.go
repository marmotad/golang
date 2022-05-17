package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//生成随机数种子
	randomNumber := rand.Intn(100)
	//定义变量并生成随机数
	var intNum int
	for i := 0; i < 5; i++ {
		fmt.Print("请输入你猜的数字：")
		fmt.Scan(&intNum)
		//输入猜的数字
		if intNum == randomNumber {
			//如果输入的数字和生成的随机数相同则打印猜对了，并将i赋值为4，结束本次循环
			fmt.Print("猜对了,")
			i = 4
		} else if intNum >= randomNumber {
			//如果输入的数字比生成的随机数大，输出提示并进行下一次循环
			fmt.Printf("太大了,你还有%d 次机会\n", 4-i)
		} else if intNum <= randomNumber {
			//如果输入的数字比生成的随机数小，输出提示并进行下一次循环
			fmt.Printf("太小了,你还有%d 次机会\n", 4-i)
		}
		if i == 4 {
			//如果i=5，询问是否再来一次
			fmt.Print("是否再来一次(Y/N)：")
			var yes string
			fmt.Scan(&yes)
			//定义yes变量，判断是否再来一次
			if yes == "y" || yes == "Y" {
				//如果输入y则再来一次
				i = 0
				//将i赋值为0，使循环重新开始
				randomNumber = rand.Intn(100)
				//重新生成随机数
			} else if yes == "n" || yes == "N" {
				//如果输入n，结束游戏
				fmt.Println("游戏结束！！！")
			}
		}
	}
}
