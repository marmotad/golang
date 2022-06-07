package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	var host string
	var verbor, help bool
	//绑定命令行与命令的关系
	flag.IntVar(&port, "p", 22, "ssh-port")
	flag.StringVar(&host, "H", "127.0.0.1", "ssh-host")
	flag.BoolVar(&verbor, "v", false, "detail log")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage flagargs [-H 127.0.0.1] [-P 22] [-v]")
		flag.PrintDefaults()
	}

	//解析命令行与命令的关系
	flag.Parse()
	if help {
		flag.Usage()
	} else if !help {
		fmt.Println(port, host, verbor)
	}
}
