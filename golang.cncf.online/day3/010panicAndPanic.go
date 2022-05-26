package main

import "fmt"

func testerr() (err error) {
	defer func() {
		//如果err值不为nil，使用recover进行错误处理，打印错误
		if e := recover(); e != nil {
			e = fmt.Errorf("%v", e)
		}
	}()
	//抛出错误
	panic("error")
	//如果有错误信息，返回错误
	return
}

//go语言提供panic和recover函数用于处理运行时错误，当调用panic抛出错误，中断原有的控制流程，常用于不可修复性错误。recover函数用于终止错误处理流程，仅在defer语句的函数中有效，用于截取错误处理流程，recover只能捕获到最后一个错误
func main() {
	/*
		//对错误进行延迟处理
		defer func() {
			//如果err值不为nil，使用recover进行错误处理，打印错误
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		fmt.Println("main start")
		panic("error")
		fmt.Println("over")
	*/
	err := testerr()
	fmt.Println(err)
}
