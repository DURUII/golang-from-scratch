package main

import (
	"example.com/pkg/series"
	"fmt"
	// 此外，还可以写 github 链接，注意 go-get 仓库不要 src 目录
)

func init() {
	fmt.Println("X")
}

func main() {
	fmt.Println(series.GetFibonacciList(5))
}
