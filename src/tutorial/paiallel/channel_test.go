package main

import (
	"fmt"
	"testing"
	"time"
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

// 有缓冲 channel 示例
// 并发地向三个镜像站点发出请求，它们分别将收到的响应发送到带缓存channel，取最快返回结果的
// mirroredQuery函数可能在另外两个响应慢的镜像站点响应之前就返回了结果
// 如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG
func TestBufferedChannel(t *testing.T) {
	request := func(s string) string { return s }

	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	println(<-responses) // return the quickest response

}

// 单方向 channel，只能发送或接收
func TestOneWayChannel(t *testing.T) {
	// 只能接收
	fmt.Println(make(<-chan int))
	// 只能发送
	fmt.Println(make(chan<- int))

	// https://gopl-zh.github.io/ch8/ch8-04.html
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestRSS(t *testing.T) {
	ticker := time.Tick(1 * time.Second)

	// 使用 for 循环从 channel 中读取时间
	for now := range ticker {
		fmt.Println("当前时间:", now)
	}
}
