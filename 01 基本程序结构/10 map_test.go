package ch01

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 声明 + 初始化： key 的类型必须支持“==”和“!=”两种比较操作符
	// 例如，函数类型、map，以及切片只支持与 nil 的比较，就不能做 key
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	// 声明
	m2 := map[int]int{}
	m2[3] = 27
	t.Logf("len m2=%d", len(m2))

	// make 方法，可以设定 cap，但不能用 cap(m3)
	m3 := make(map[string]int, 10)
	t.Log(m3, len(m3))
}

func TestMapBasedSet(t *testing.T) {
	// Go 语言没有提供 set，但可以通过 map 实现类似的功能
	m1 := map[int]int{2: 8}

	/*```python
	map = {2:8}

	In [9]: map[3]
	KeyError: 3

	In [12]: type(map.get(3))
	Out[12]: NoneType
	```*/

	/*```java
	import java.util.*;

	public class MyClass {
	  public static void main(String args[]) {
		Map<Integer, Integer> map = new HashMap<>();
		// null
		System.out.println(map.get(4));
	  }
	}
	```*/

	// 当 key 不存在时，默认返回零值（与 Python/Java 不同）
	// 所以我们无法通过 val 判断出，究竟是因为 key 不存在返回的零值，
	// 还是因为 key 本身对应的 value 就是 0
	if val, ok := m1[3]; ok { // “comma ok”的惯用法，这类似 Python 中 map.get(key, default)，而 default 系统自动返回
		t.Log("retrieved val =", val)
	} else {
		t.Log("key not found")
	}
}

func TestMapOperation(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	// 插入数据
	m1[5] = 25
	t.Log(m1[2], len(m1)) // 不支持 cap

	// 删除元素
	delete(m1, 4)
	delete(m1, 100) // 执行不会失败，不会抛出运行时的异常
	t.Log(len(m1))
}

func TestMapIterationOrder(t *testing.T) {
	// 注意，map 输出次序是有意而为地随机化的，用于防止程序员偷懒
	// *判断 map 中是否存在某个值，需要遍历 map
	m := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	for k := 0; k < 10; k++ {
		// 也可以写成 key, _ := range m
		for key := range m {
			fmt.Printf("%d ", key)
		}
		fmt.Println()
	}
}

func TestMapWithFuncValue(t *testing.T) {
	// 与 Python/JavaScript 类似，支持部分函数式编程特性
	intFunc := map[string]func(op int) int{}
	intFunc["cubic"] = func(op int) int { return op * op * op }
	t.Log(intFunc["cubic"](9))
}
