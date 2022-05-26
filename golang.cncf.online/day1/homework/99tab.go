package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				fmt.Println()
				break
			} else if i >= j {
				fmt.Print(j, "*", i, "=", j*i, "\t")
			}
		}
	}
}
