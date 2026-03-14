package main

import (
	"fmt"
	"time"
)

// 演示 Go 语言的 Channel（通道）

func main() {
	// ========== 创建通道 ==========
	// 无缓冲通道
	ch1 := make(chan int)

	// 有缓冲通道
	ch2 := make(chan string, 3)

	// ========== 发送和接收 ==========
	// 发送数据（阻塞操作）
	go func() {
		ch1 <- 42 // 发送数据
	}()

	// 接收数据（阻塞操作）
	val := <-ch1
	fmt.Printf("接收到数据: %d\n", val)

	// 带缓冲的通道
	ch2 <- "hello"
	ch2 <- "world"
	ch2 <- "!"

	fmt.Printf("缓冲通道内容: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", <-ch2)
	}
	fmt.Println()

	// ========== 关闭通道 ==========
	// 发送方应该关闭通道
	ch3 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch3 <- i
	}
	close(ch3) // 关闭通道

	// 接收关闭的通道
	fmt.print("接收关闭通道的值: ")
	for {
		val, ok := <-ch3
		if !ok {
			break // 通道已关闭
		}
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// ========== 单向通道 ==========
	// 只发送通道
	sendOnly := make(chan<- int)
	// sendOnly <- 1 // 可以发送
	// <-sendOnly    // 不能接收

	// 只接收通道
	recvOnly := make(<-chan int)
	// recvOnly <- 1 // 不能发送
	// <-recvOnly    // 可以接收

	// 函数参数使用单向通道
	producer := func(ch chan<- int) {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("消费: %d\n", v)
		}
	}

	// 使用单向通道
	ch4 := make(chan int)
	go producer(ch4)
	consumer(ch4)

	// ========== 通道作为函数返回值 ==========
	resultCh := generate()
	for v := range resultCh {
		fmt.Printf("生成: %d\n", v)
	}

	fmt.Println("主函数结束")
}

// 返回通道的函数
func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return ch
}
