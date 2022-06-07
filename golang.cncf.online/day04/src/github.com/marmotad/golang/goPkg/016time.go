package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	//打印now变量的类型和值
	fmt.Printf("%T \n %v \n", now, now)
	//指定输出格式time.Now.Format()
	// 2000/01/02 08:10:01
	// 2006 年
	// 01 月
	// 02 天
	// 24进制小时 15
	// 分钟 04
	// 秒 05
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//转换Unix时间戳
	fmt.Println(time.Now().Unix())
	//转换unix时间戳（纳秒为单位）
	fmt.Println(time.Now().UnixNano())
	//将时间转换为字符串格式
	time1, err := time.Parse("2006-01-02 15:04:05", "2022-01-02 15:04:05")
	fmt.Println(time1, err)
	//将unix时间转换为,sec:秒，nsec:纳秒
	t := time.Unix(1, 0)
	fmt.Println(t)
	//计算两个时间的时间差
	time2 := time.Unix(0, 0).Sub((time.Now()))
	fmt.Printf("%T, %v\n", time2, time2)

	//时间区间
	// time.Second
	// time.Minute
	// time.Hour
	// 3h2m4s
	//打印当前时间
	fmt.Println(time.Now())
	//让程序sleep 5s
	time.Sleep(time.Second * 5)
	//打印sleep后的时间
	fmt.Println(time.Now())
	//获取当前时间的三小时前的时间
	fmt.Println(time.Now().Add(-3 * time.Hour))
	//将字符串转换为时间区间
	date, err := time.ParseDuration("-3h2m4s")
	fmt.Println(err, date)
	//将时间转换为小时
	fmt.Println(date.Hours())
	//将时间转换为分钟
	fmt.Println(date.Minutes())
	//将时间转换为秒
	fmt.Println(date.Seconds())
}
