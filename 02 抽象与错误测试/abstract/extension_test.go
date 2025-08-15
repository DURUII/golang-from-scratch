package abstract

import (
	"fmt"
	"testing"
)

// 父类
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	// Go 是不⽀持继承的
	fmt.Println(" ", host)
}

// 子类
type Dog struct {
	Pet // 组合复用，无继承；不支持方法覆盖
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}

func TestStructOverride(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Jack") // ...  Jack
}
