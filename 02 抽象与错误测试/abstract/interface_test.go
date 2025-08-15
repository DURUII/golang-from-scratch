package abstract

import "testing"

// 接口是对行为的抽象，尽量定义小接口，方法个数在 1~3 个之间
// 非入侵性：接口的实现不需要写 implement，不与接口定义依赖绑定 （与 Java 不同）
type GoProgrammer struct {
}

// Duck Type 你和 Programmer 方法签名一样，那你就是 Programmer
// 甚至你可以先写 GoProgrammer 后写 interface
func (p *GoProgrammer) WriteHelloWorld() string {
	return "Hello World, Go!"
}

// 接口用来定义对象间的“交互协议”，接口定义可以写在接口使用包里 （与 Java 不同）
type Programmer interface {
	WriteHelloWorld() string
}

// 这种写法常用来在编译时确认接口实现
var _ Programmer = &GoProgrammer{}

func TestClient(t *testing.T) {
	// Java 风格的新建也是抽象化类 = 实现类，支持多态
	var prog Programmer = new(GoProgrammer)
	t.Log(prog.WriteHelloWorld())
}
