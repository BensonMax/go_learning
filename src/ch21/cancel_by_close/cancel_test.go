package concurrency

import (
	"fmt"
	"testing"
	"time"
)

//对channel 方法，判断
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true //如果是cancelChan 返回true
	default:
		return false //默认返回 false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} //将一个空结构 写入Struct{}，传递给 cancelChan
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan) //调用广播方法结束 cancelChan
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)  //输出 4 Cancelled
	cancel_2(cancelChan)
	/*
		输出
		1 Cancelled
		4 Cancelled
		0 Cancelled
		2 Cancelled
		3 Cancelled
	*/
	time.Sleep(time.Second * 1)
}
