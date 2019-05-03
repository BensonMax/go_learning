package polymorphism

import (
	"fmt"
	"testing"
)

//"多态"

type Code string //别名（自定义类型）

//定义接口
type Programmer interface {
	WriteHelloWorld() Code
}

//定义数据结构
type GoProgrammer struct {
}

//接口实现
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

//定义数据结构
type JavaProgrammer struct {
}

//接口实现
func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

//对接口编程
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) //打印
}

//测试程序
func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)     //获得一个GoProgrammer指针，赋与goProg
	javaProg := new(JavaProgrammer) //获得一个JavaProgrammer指针，赋与goProg
	writeFirstProgram(goProg)       //对goProg调用writeFirstProgram方法
	writeFirstProgram(javaProg)     //对javaProg调用writeFirstProgram方法
}
