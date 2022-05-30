package main

import (
	"fmt"
	"os"
)

func main () {
	args := os.Args

	var str string

	if len(args) > 0 {
		for _, arg  := range args[1:] {
			str = str + " " + arg
		}
		fmt.Println(str)
	}
}