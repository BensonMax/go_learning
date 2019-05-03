package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//定义异步服务，返回一个string channel
func AsyncService() chan string {
	retCh := make(chan string) //非buffer channel 会被阻塞
	//retCh := make(chan string, 1)  //buffer channel 不会被阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

func AsyncService2() chan string {
	retCh2 := make(chan string, 1)
	//retCh := make(chan string, 1)  //buffer channel 不会被阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh2 <- ret
		fmt.Println("service exited.")
	}()
	return retCh2
}

func TestAsyncService2(t *testing.T) {
	retCh2 := AsyncService2()
	otherTask()
	fmt.Println(<-retCh2)
	time.Sleep(time.Second * 2)
}
