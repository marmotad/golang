package main

import "log"

func main() {
	//打印日志
	log.Printf("使用printf打印日志1")
	//使用SetPrefix设置日志前缀通常用于日志分类
	log.SetPrefix("prefix:")
	log.Printf("使用printf打印日志2")
	//设置消息格式
	//打印日志文件名(相对路径)
	//log.SetFlags(log.Flags() | log.Lshortfile)
	log.Printf("使用printf打印日志3")
	//打印日志文件名(绝对路径)
	log.SetFlags(log.Flags() | log.Llongfile)
	log.Printf("使用printf打印日志4")
	//获取当前日志的标签
	//log.Panic打印错误，默认会结束运行
	//log.Panic("打印Panic")
	//Fatalf打印日志，打印完成后会退出
	log.Fatalf("使用Fatalf打印日志1")
	log.Fatalf("使用Fatalf打印日志2")
	log.Printf("使用printf打印日志")
}
