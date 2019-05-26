package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

/*
并发任务：只执行一次   sync.Once
*/
type Singleton struct {
	data string
}

var singleInstance *Singleton
var once sync.Once

//func GetSingletonObj() *Singleton {
//	once.Do(func() {
//		fmt.Println("Create Obj")
//		singleInstance = new(Singleton)
//	})
//	return singleInstance
//}
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()                //通过GetSingletonObj方法获取对象，赋值给obj
			fmt.Printf("%X\n", unsafe.Pointer(obj)) //打印obj 对象的地址
			wg.Done()                               //线程组结束
		}()
	}
	wg.Wait()
}
