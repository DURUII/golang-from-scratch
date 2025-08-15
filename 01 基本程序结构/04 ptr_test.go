package ch01

import "testing"

func TestPointer(t *testing.T) {
	// 不支持指针运算，限制危险功能（杜绝野指针、内存越界等低级错误）
	var a = 1
	var aPtr = &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, &a)
}
