package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	//通过select 实现多路选择和超市控制
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100): //时间控制，当时间超过100ms 后自动执行下一语句
		t.Error("time out")
	}
}
