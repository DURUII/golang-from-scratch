package abstract

import (
	"fmt"
	"testing"
	"unsafe"
)

// 对数据的封装/抽象（与 C 类似），没有逗号
// Go 用标识符名称的首字母大小写来判定这个标识符是否为导出标识符
// *用空标识符“_”作为结构体类型定义中的字段名称：占位与内存对齐
// *空结构体类型变量的内存占用为 0
// *支持嵌入字段
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 对行为的定义，不写在 struct 里，不是典型的 OOP
// 指针接收器，避免内存分配与复制；结构体接收器，无法修改对应原对象字段
func (e *Employee) String() string {
	// 没有对象复制产生
	fmt.Println("Address = ", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("{id:%s name:%s age:%d}", e.Id, e.Name, e.Age)
}

// Go 语言不支持继承
func TestCreateData(t *testing.T) {
	e := Employee{"1", "Jack", 20}
	e2 := new(Employee) // 返回指针类型
	e2.Id = "123"
	e2.Name = "Tom"
	e2.Age = 30
	// ch11.Employee, *ch11.Employee
	t.Logf("%T", e)
	t.Logf("%T", e2)
	// 不需要箭头符号（与 C 不同）
	t.Log(e2.String())
	fmt.Println("Address = ", unsafe.Pointer(&e2.Name))
	// 如果结构体的所有成员变量都是可比较的，那么结构体就可比较
	fmt.Println(*e2 == e)
}

// 变量的内存地址值必须是其类型本身大小的整数倍
// 对于结构体而言，它最长字段长度与系统对齐系数两者之间较小的那个的整数倍
type Q struct {
	b byte // 1
	// padding 7
	i int64  // 8
	u uint16 // 2
	// padding 6
}

type S struct {
	b byte // 1
	// padding 1
	u uint16 // 2
	// padding 4
	i int64 // 8
}

func TestStructSize(t *testing.T) {
	var q Q
	fmt.Println(unsafe.Sizeof(q)) // 24
	var s S
	fmt.Println(unsafe.Sizeof(s)) // 16
	// 没有构造析构函数的概念，所以所有字段都是对应的零值
	fmt.Println(s, q)
}
