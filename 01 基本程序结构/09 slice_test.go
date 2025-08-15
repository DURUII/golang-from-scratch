package ch01

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefSlice(t *testing.T) {
	a := []int{1, 2, 3}
	fn := func(x []int) { x[1] = 4 }
	fn(a)
	t.Log(a)
}

func TestSliceCap(t *testing.T) {
	// 没有...表示，一个切片（可变长数组），支持零值可用*
	// 注意，map 不支持零值可用
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 切片有len/cap，用 make 指定容量，避免扩容复制*
	// *这里代表，有 3 个 0 元素，容量为 4
	var s1 = make([]int, 3, 4)
	// 当存储空间不足，自动开辟新的存储空间，并拷贝原有数值
	// 在循环操作中频繁 append 自动扩容会放大性能损失，减慢程序的运行
	s1 = append(s1, 10)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 40)
	t.Log(s1, len(s1), cap(s1))
}

func TestSliceSharedMem(t *testing.T) {
	/*
		每个切片包含三个字段：
		array: 是指向底层数组的指针；
		len: 是切片的长度，即切片中当前元素的个数；
		cap: 是底层数组的长度，也是切片的最大容量，cap 值永远大于等于 len 值（动态扩容）。
	*/
	year := []string{
		"Jan", "Feb", "Mar", "Apr", "May",
		"Jun", "Jul", "Aug", "Sep", "Oct", "Nov",
		"Dec",
	}
	// 可以用 array[low : high : max] 基于一个已存在的数组创建切片
	// len = high - low, cap = max - low
	Q2 := year[3:6]
	// 这里的 cap 是 9，剩余连续存储空间
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	summer[0] = "Unknown"
	// 共享存储空间
	t.Log(year, Q2, summer)
}

func TestSliceSharedMem2(t *testing.T) {
	transform := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			arr[i] = arr[i] * 2
		}
	}

	nums := [...]int{1, 2, 3, 4, 5}
	transform(nums[:])
	fmt.Println(nums)
	assert.Equal(t, nums, [...]int{2, 4, 6, 8, 10})
}

// *Fix for-range Issue Again in Go 1.22
// 要复现这个错误，需要在 go.mod 文件中设置 go 1.22
func TestSliceForRangeBug(t *testing.T) {
	a := []int{1, 2, 3}
	b := []*int{}
	// 注意，遍历过程没有返回集合中的实际元素
	// 而是复制在了一个“固定”的变量（值传递），例如 i
	// *复制操作可能会存在性能问题
	for _, i := range a {
		b = append(b, &i)
	}
	// 相当于
	//var i int
	//for k := 0; k < len(a); k++ {
	//	i = a[k]
	//	b = append(b, &i)
	//}

	for _, j := range b {
		fmt.Print(*j, " ")
	}
}
