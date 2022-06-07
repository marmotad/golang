package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/marmotad/golang/gopkg"
)

func main() {
	beego.Run()
	fmt.Println(gopkg.VERSION)
}
