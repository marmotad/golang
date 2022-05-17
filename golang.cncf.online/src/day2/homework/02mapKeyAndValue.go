package main

import (
	"fmt"
)

func main() {
	var keySlice, valueSlice []string
	map1 := map[string]string{"姓名": "lisi"}
	for k, v := range map1 {
		keySlice = append(keySlice, k)
		valueSlice = append(valueSlice, v)
	}
	fmt.Println(keySlice, valueSlice)

}
