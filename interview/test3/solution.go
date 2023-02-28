package solution

import (
	"fmt"
)

func doTask(c1,c2 chan int) {
	for {
		select {
		case task,ok := <-c1 :
			if !ok {
				return
			}
			fmt.Printf("进程%v计算\n",task)
			v := <- c2
			v += task
			fmt.Printf("当前总和%v\n",v)
			c2 <- v
		}
	}
}

func solution() {
	var n = 10
	var c1 = make(chan int,n)
	for i := 1 ; i <= 10 ; i ++ {
		c1 <- i
	}
	close(c1)
	var c2 = make(chan int,1)
	c2 <- 0
	for i := 0 ; i < 3 ; i ++ {
		go doTask(c1,c2)
	}
}