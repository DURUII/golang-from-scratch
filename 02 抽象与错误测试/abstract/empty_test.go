package abstract

import (
	"fmt"
	"testing"
)

// 类似于 C 中的 void* 和 Java 中的 Object
func DoSomething(p interface{}) {
	// 类型 switch 专用语法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(1)
	DoSomething("1")
	DoSomething(fmt.Println)

	// 再看 fmt.Println
	// type any = interface{}
	// func Println(a ...any) (n int, err error) {
	//	return Fprintln(os.Stdout, a...)
	// }
	fmt.Println(fmt.Println("Hello World Again"))
}
