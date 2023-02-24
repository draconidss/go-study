package main

import (
	"fmt"
	"testing"
	"time"
)

//手动触发宕机
func TestPanic(t *testing.T) {
	panic("crash")
}

func setData(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}

//捕获异常，恢复程序或做收尾工作
//recover 调用后，抛出的 panic 将会在此处终结，不会再外抛，
//但是 recover，并不能任意使用，它有强制要求，必须得在 defer 下才能发挥用途。
func TestRecover(t *testing.T) {
	setData(20)

	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
}

//测试panic无法跨协程
func TestCoroutine(t *testing.T) {
	// 这个 defer 并不会执行
	defer fmt.Println("in main")

	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(2 * time.Second)
}
