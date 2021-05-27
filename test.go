package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*8)
	defer cancel() // 防止任务比超时时间短导致资源未释放
	// 启动协程
	go task(ctx)
	// 主协程需要等待，否则直接退出
	time.Sleep(time.Second * 10)
}

func task(ctx context.Context) {
	ch := make(chan struct{}, 0)
	// 真正的任务协程
	go func() {
		// 模拟10秒耗时任务 这里修改成大于4秒
		time.Sleep(time.Second * 5)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
