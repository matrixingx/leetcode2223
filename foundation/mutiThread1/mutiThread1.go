package mutithread1

import (
	"context"
	"fmt"
	"time"
)

// 三个线程交替输出1-i
func printI(n int) {
	ctx := context.Background()
	ctx,cancel := context.WithCancel(ctx)
	ch1,ch2,ch3 := make(chan int),make(chan int),make(chan int)
	go func(ctx context.Context,inputCh,outputCh chan int,i,n int,cancel context.CancelFunc) {
		thread(ctx,inputCh,outputCh,i,n,cancel)
	}(ctx,ch1,ch2,1,n,cancel)
	go func(ctx context.Context,inputCh,outputCh chan int,i,n int,cancel context.CancelFunc) {
		thread(ctx,inputCh,outputCh,i,n,cancel)
	}(ctx,ch2,ch3,2,n,cancel)
	go func(ctx context.Context,inputCh,outputCh chan int,i,n int,cancel context.CancelFunc) {
		thread(ctx,inputCh,outputCh,i,n,cancel)
	}(ctx,ch3,ch1,3,n,cancel)
	ch1 <- 1
	select {
	case <- ctx.Done():
		time.Sleep(100 * time.Millisecond)
		return
	}
}

func thread(ctx context.Context,inputCh,outputCh chan int,i,n int,cancel context.CancelFunc) {
	for {
		select {
		case v := <- inputCh:
			if v > n {
				fmt.Printf("%v thread return\n",i)
				cancel()
				return
			}
			fmt.Printf("%v thread print %v\n",i,v)
			outputCh <- v+1
		case <- ctx.Done():
			fmt.Printf("%v thread return\n",i)
			return
		}
	}
}
