package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

func test3() {
	threadNum := 5    //1. 五个信号量
	wg.Add(threadNum) //2.开启5个线程
	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}
	//3.等待子线程结束
	wg.Wait()
	fmt.Println(counter)
}

func incCounter(index int) {
	defer wg.Done()
	spinNum := 0
	for {
		//2.1原子操作
		old := counter
		//在 Go 语言中，`atomic.CompareAndSwapInt32` 是一个原子操作函数，用于在无需锁的情况下，安全地对共享变量进行操作。这个函数是 `sync/atomic` 包中提供的几个原子操作之一，主要用于并发环境中保证操作的原子性。
		//`atomic.CompareAndSwapInt32` 函数的签名如下：
		//func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
		//参数说明：
		//- `addr` 是指向要操作的 32 位整数的指针。
		//- `old` 是预期的旧值。
		//- `new` 是想要设置的新值。
		//
		//该函数的工作机制如下：
		//- 如果 `addr` 指向的值与 `old` 给定的值相同，它将该值更新为 `new` 并返回 `true`，表示交换成功。
		//- 如果 `addr` 指向的值与 `old` 不同，它将不做任何操作并返回 `false`，表示交换失败。
		//
		//这个函数是原子的，即在多线程环境中，这个操作不会与其他操作交错执行，可以保证不会出现数据竞争（race condition）的问题。
		//
		//使用场景：
		//- 在多线程环境中，当需要更新一个共享的 32 位整数变量，并且希望这个更新操作是原子的，以避免并发问题时，可以使用 `CompareAndSwapInt32`。
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)
}

func test() {
	// 创建一个缓冲channel，缓冲大小为10，表示同时只允许10个goroutine运行
	semaphore := make(chan struct{}, 10)

	var wg sync.WaitGroup
	const totalRequests = 100 // 总请求数

	for i := 0; i < totalRequests; i++ {
		wg.Add(1) // 增加WaitGroup的计数

		// 使用Go语言的goroutine来处理请求
		go func(requestID int) {
			defer wg.Done() // 请求处理完毕后通知WaitGroup

			// 尝试获取一个信号量，如果无法获取，则阻塞等待
			semaphore <- struct{}{}

			// 模拟请求处理
			fmt.Printf("Request %d is being processed\n", requestID)
			// 假设请求处理需要一些时间
			time.Sleep(time.Second)

			// 请求处理完毕，释放信号量
			<-semaphore
		}(i)
	}

	// 等待所有请求处理完毕
	wg.Wait()
	fmt.Println("All requests have been processed.")
}

func test1() {
	var (
		ch = make(chan struct{}, 10)
		wg sync.WaitGroup
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			ch <- struct{}{}

			// 模拟请求处理
			fmt.Printf("Request %d is being processed\n", i)
			// 假设请求处理需要一些时间
			time.Sleep(time.Second)

			<-ch
		}(i)
	}
	wg.Wait()
	fmt.Println("All requests have been processed.")
}
