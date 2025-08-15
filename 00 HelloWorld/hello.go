// package 名不需要与目录名一致
// 但同一目录的 Go 代码 package 要保持一致
// 非测试程序入口必须是 main 包 + main 函数，文件名不一定是 main.go
package main

import (
	"fmt" // 依赖包的路径（默认在标准库目录下）
	"os"
)

// main 函数并不一定是第一个被执行的函数，Go 语言会先初始化常量变量，并自动调用包初始化的 init 函数
func main() {
	fmt.Print("Hello World!")
	// main 函数不支持入参
	if len(os.Args) > 1 { // 极简的 Go 语言 不需要括号、分号
		// 大写代表包外可以访问
		fmt.Println(os.Args[1])
	}
	// main 函数不支持返回值（与 C 不同）
	os.Exit(0)
}
