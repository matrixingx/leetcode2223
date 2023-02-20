package test1

import "fmt"

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	endChan := make(chan int)
	go thread(ch1, ch2, endChan,"a")
	go thread(ch2, ch3, endChan,"b")
	go thread(ch3, ch1, endChan,"c")
	ch1 <- 0
	select {
	case <-endChan:
		return
	}
}

func thread(inputChan,outputChan, endChan chan int, s string) {
	for {
		select {
		case v := <-inputChan:
			if v >= 300 {
				endChan <- v
				return
			}
			fmt.Println(s)
			outputChan <- v + 1
		case <-endChan:
			return
		}
	}
}