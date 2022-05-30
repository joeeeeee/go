package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 根据输入的文件名来查询重复的文字
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		file, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Printf("error ")
		}

		row := strings.Split(string(file), "\n")

		for _, r := range row {
			counts[(r)] ++
		}
	}

	fmt.Println("==========================")
	for key, count := range counts {
		if count > 1 {
			fmt.Println("repeat :", count, "key", key)
		}
	}

}
