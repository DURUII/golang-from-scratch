package ch00

import (
	"bytes"
	"fmt"
	"reflect"
	"slices"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func shareUnderlyingArray(s1, s2 []int) bool {
	//return (*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data == (*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data
	return unsafe.SliceData(s1) == unsafe.SliceData(s2)
}

// https://github.com/golang/go/blob/master/src/runtime/slice.go#L13
func TestSliceShareUnderlyingArray(t *testing.T) {
	a := make([]int, 32)
	b := a[:16] // 切片共享底层数组，b: len: 15, cap: 31
	assert.Equal(t, true, shareUnderlyingArray(a, b))
	a = append(a, 1) // 这个操作会让 a 扩容，重新分配内存
	a[2] = 42
	assert.Equal(t, 64, cap(a))                        // ok
	assert.Equal(t, false, shareUnderlyingArray(a, b)) // ok
	assert.Equal(t, a[2], b[2])                        // fail
}

func TestSliceShareUnderlyingArray2(t *testing.T) {
	path := []byte("bbs.wps.cn/node/22")
	firstSepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:firstSepIndex]
	dir2 := path[firstSepIndex+1:]
	assert.Equal(t, "bbs.wps.cn", string(dir1))
	assert.Equal(t, "node/22", string(dir2))

	dir1 = append(dir1, "/topics"...)
	// 共享内存，对 dir1 的 append 操作没有超出 cap，于是数据扩展到了 dir2
	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	path2 := []byte("bbs.wps.cn/node/22")
	firstSepIndex2 := bytes.IndexByte(path2, '/')
	dir3 := path2[:firstSepIndex2:firstSepIndex2] // Limited Capacity
	dir4 := path2[firstSepIndex2+1:]

	dir3 = append(dir3, "/topics"...)
	// 重新分配内存
	fmt.Println("dir3:", string(dir3))
	fmt.Println("dir4:", string(dir4))
}

func TestSliceEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	assert.Equal(t, true, slices.Equal(s1, s2))      // 先比较长度 再逐一比较元素
	assert.Equal(t, true, reflect.DeepEqual(s1, s2)) // 通用函数
}
