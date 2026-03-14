# Go Channel Demo

## 简介

演示 Go 语言中的通道（Channel）创建、发送接收和关闭。

## 基本原理

Channel 是 Goroutine 之间的通信机制，支持阻塞发送和接收。可以有缓冲或无缓冲。

## 启动和使用

### 环境要求
- Go 1.21+

### 安装和运行

```bash
cd go-channel-demo
go run main.go
```

## 教程

### 创建通道

```go
// 无缓冲通道
ch := make(chan int)

// 有缓冲通道
ch := make(chan string, 3)
```

### 发送接收

```go
// 发送
ch <- value

// 接收
val := <-ch
```

### 关闭通道

```go
close(ch)

// 接收时检查是否关闭
val, ok := <-ch
if !ok {
    // 通道已关闭
}
```

### 单向通道

```go
// 只发送
ch := make(chan<- int)

// 只接收
ch := make(<-chan int)
```

### 要点

1. 通道默认阻塞，直到有接收者
2. 关闭后不能发送但可以接收
3. 单向通道用于函数参数
4. 通道是引用类型
