package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): //如果收到context结束通知，return true
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	//ctx, cancel := context.WithCancel(context.Background())
	//for i := 0; i < 5; i++ {
	//	go func(i int, ctx context.Context) {
	//		for {
	//			if isCancelled(ctx) {
	//				break
	//			}
	//			time.Sleep(time.Millisecond * 5)
	//		}
	//		fmt.Println(i, "Cancelled")
	//	}(i, ctx)
	//}
	//cancel()
	//time.Sleep(time.Second * 1)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
					time.Sleep(time.Millisecond * 5)
				}
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
