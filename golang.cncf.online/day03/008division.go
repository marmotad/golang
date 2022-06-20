package main

import (
	"errors"
	"fmt"
)

// 返回值 怎么定义错误类型 error
// 怎么创建错误类型的值, errors.New(), fmt.Errorf()
// 无错误 nil

//定义一个函数，返回值是int和error
func division(a, b int) (int, error) {
	//如果被除数为0，返回-1
	if b == 0 {
		return -1, errors.New("division is zero")
	}
	//如果被除数不为0，返回结果和error为nil
	return a / b, nil
}

func main() {
	fmt.Println(division(10, 2))
	//如果err值为nil，打印运算结果
	if v, err := division(6, 0); err == nil {
		fmt.Println(v)
		//如果err值不为nil，打印error
	} else if err != nil {
		fmt.Println(err)
	}

	e := fmt.Errorf("Error: %s", "division by zero")
	fmt.Printf("%T, %v\n", e, e)
}
