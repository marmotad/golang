package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串=>其他类型
	//=> bool
	if v, err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v)
	} else if err != nil {
		fmt.Println(err)
	}
	// 字符串转int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	} else if err != nil {
		fmt.Println(err)
	}
	/*
		参数1 数字的字符串形式
		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	*/
	if v, err := strconv.ParseInt("64", 16, 64); err == nil {
		fmt.Println(v)
	} else if err != nil {
		fmt.Println(err)
	}
	//float
	// ParseFloat 将字符串转换为浮点数
	// s：要转换的字符串
	// bitSize：指定浮点类型（32:float32、64:float64）
	// 如果 s 是合法的格式，而且接近一个浮点值，
	// 则返回浮点数的四舍五入值（依据 IEEE754 的四舍五入标准）
	// 如果 s 不是合法的格式，则返回“语法错误”
	// 如果转换结果超出 bitSize 范围，则返回“超出范围”
	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)          // 0.12345679104328156
	fmt.Println(float32(f), err) // 0.12345679
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err) // 0.12345678901234568

	//Sprintf:来格式化字符串
	sd := fmt.Sprintf("%d", 12)
	fmt.Println(sd)

	sf := fmt.Sprintf("%.3f", 10.99)
	fmt.Println(sf)
	//FormatBool根据b的值返回"true"或"false"
	fmt.Printf("%q\n", strconv.FormatBool(false))
	//int转字符串
	fmt.Printf("%q\n", strconv.Itoa(12))
	//FormatInt 返回 i 在給定基數中的字符串表示形式，即 2 <= base <= 36。結果使用小寫字母 'a' 到 'z' 表示數字值 >= 10。
	fmt.Printf("%q\n", strconv.FormatInt(12, 16))
	//根据格式 fmt 和精度 prec 将浮点数 f 转换为字符串。假设原始值是从 bitSize 位的浮点值(32 表示 float32，64 表示 float64)获得的，它会对结果进行四舍五入
	fmt.Printf("%q\n", strconv.FormatFloat(10.1, 'f', -1, 64))
}
