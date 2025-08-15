package ch01

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStringDefaultValue(t *testing.T) {
	// 字符串是原生数值类型，的默认值是""，而不是 None/nil
	var str string
	// 获取 Go 字符串长度操作的时间复杂度是 O(1)
	// string 只能和 string 做拼接，Go 不会做类型转换（如自动加上 .toString()）
	t.Log("*"+str+"*", len(str))
}

func TestRawStringLiteral(t *testing.T) {
	// Figlet
	// “所见即所得”的原始字符串 （类似 Python 的 """ <content> """）
	s := `
     _____          ___           ___           ___                 
    /  /::\        /__/\         /  /\         /__/\        ___     
   /  /:/\:\       \  \:\       /  /::\        \  \:\      /  /\    
  /  /:/  \:\       \  \:\     /  /:/\:\        \  \:\    /  /:/    
 /__/:/ \__\:|  ___  \  \:\   /  /:/~/:/    ___  \  \:\  /__/::\    
 \  \:\ /  /:/ /__/\  \__\:\ /__/:/ /:/___ /__/\  \__\:\ \__\/\:\__ 
  \  \:\  /:/  \  \:\ /  /:/ \  \:\/:::::/ \  \:\ /  /:/    \  \:\/\
   \  \:\/:/    \  \:\  /:/   \  \::/~~~~   \  \:\  /:/      \__\::/
    \  \::/      \  \:\/:/     \  \:\        \  \:\/:/       /__/:/ 
     \__\/        \  \::/       \  \:\        \  \::/        \__\/  
                   \__\/         \__\/         \__\/
`
	fmt.Println(s)
}

func TestStringByteLength(t *testing.T) {
	// string 不是引用/指针类型，而是原生支持字符串，空值为 ""
	/*
		```python
		'李'.encode('utf-8').hex()
		bytes.fromhex('e69d8e').decode('utf-8')
		```
	*/
	// string 是 `只读` 的 byte 切片
	// 因此 len 返回 byte 数
	var s = "\xe6\x9d\x9c\xe7\x9d\xbf"
	t.Log(s, len(s))

	chars := []rune(s)
	t.Logf("杜 unicode %x", chars[0])
	t.Logf("杜 utf-8 %x", s[:3])
	// string 的 byte 数组可以存放任何数据
}

func TestStringImmutability(t *testing.T) {
	str := "hello"
	// Go 语言规定，字符串类型的值在它的生命周期内是不可改变的，这提高了字符串的并发安全性和存储利用率。
	// str[0] = 'x'
	fmt.Println(str)
}

func TestUnicodeIteration(t *testing.T) {
	s := "hello, 中国🀄!"
	// *Go 语言中的字符串值是一个可空的字节序列，也是一个可空的字符序列
	// rune 这个类型本质上是 int32，表示一个 Unicode 码点，一个 rune 实例就是一个 Unicode 字符
	fmt.Println(len(s), utf8.RuneCountInString(s))

	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringPkg(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(strings.Join(parts, "->"))
	t.Log(strings.ContainsRune(s, ','))
	// 拼写历史来由：Integer to ASCII
	s = strconv.Itoa(10)
	t.Log(string('*') + s + "*") // 注意强制类型转换
	if i, err := strconv.Atoi("100"); err == nil {
		t.Log(100 + i)
	}
}
