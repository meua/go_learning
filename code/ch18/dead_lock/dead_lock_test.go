package dead_lock

import (
	"fmt"
	"testing"
)

func TestDeadLock(t *testing.T) {
	ch := make(chan int)
	<-ch
	t.Log("你等不到我输出")
	close(ch)
}

func TestProdutorConsumer(t *testing.T) {
	// 创建一个容量为4的channel
	ch := make(chan int, 4)
	// 创建4个协程，作为生产者
	for i := 0; i < 4; i++ {
		go func() {
			ch <- 7
		}()
	}
	// 创建4个协程，作为消费者
	for i := 0; i < 4; i++ {
		go func() {
			o := <-ch
			fmt.Println("received:", o)
		}()
	}
}
