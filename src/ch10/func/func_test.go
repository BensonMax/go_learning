package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 方法名returnMultiValues 无入参，返回2个返回值，值的类型是int
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/*
输入是一个函数类型，返回是一个函数类型

func timeSpent(inner func(op int) int --> 函数类型) func(op int) int --> 函数类型

*/
func timeSpent(inner func(op int) int) func(op int) int {
	//返回一个函数
	return func(n int) int {
		start := time.Now()
		ret := inner(n)                                         //调用传入的函数
		fmt.Println("time spent:", time.Since(start).Seconds()) //返回耗时
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
