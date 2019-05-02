package interface_test

import "testing"

//定义Programmer接口
type Programmer interface {
	WriteHelloWorld() string
}

//定义GoProgrammer结构体
type GoProgrammer struct {
}

//接口实现
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
