package palallel

import (
	"fmt"
	"testing"
)

// fatal error: all goroutines are asleep - deadlock!
func TestChannelDeadLock(t *testing.T) {
	ch := make(chan int)
	<-ch
}

func TestChannelOperation(t *testing.T) {
	// 创建
	// 无缓冲
	ch1 := make(chan struct{}) // 等同于 make(chan int, 0)
	// 有缓冲
	ch2 := make(chan struct{}, 5)

	// 发送
	ch1 <- struct{}{}
	ch2 <- struct{}{}
	ch2 <- struct{}{}

	// 接收
	<-ch1       // 接收并将取出的值丢弃
	x1 := <-ch2 // 接收并将取出的值赋值给变量x
	fmt.Println(x1)
	x2, ok := <-ch2 // 接收并将取出的值赋值给变量x，同时判断通道是否关闭或为空
	fmt.Println(x2, ok)
}
