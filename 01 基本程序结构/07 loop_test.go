package ch01

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	// Go 坚持一件事情仅有一种做法的理念，对于循环只支持 for
	// 与 C++/Java 不同，不需要括号

	/*```python
	for i in range(5):
		print(i)
	```*/

	/*```java
	class Main {
		public static void main(String[] args) {
			for (int i = 0; i < 5; i++){
				System.out.println(i);
			}
		}
	}
	```*/

	// 为保持简洁，Go 语言将规定自增只能是语句，而不是表达式
	// 不支持 ++i 或者 --i 操作
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 不支持链式赋值（与 C 不同）
	//i := 10
	//x := i++
}

func TestLabelBreak(t *testing.T) {
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven = -1

	// 不带 label 的 break 语句中断执行并跳出的，
	// 是同一函数内 break 语句所在的最内层的 for、switch 或 select
loop:
	for i := 0; i < len(sl); i++ {
		switch sl[i] % 2 {
		case 0:
			firstEven = sl[i]
			break loop
		case 1:
			// do nothing
		}
	}

	println(firstEven) // 6
}
