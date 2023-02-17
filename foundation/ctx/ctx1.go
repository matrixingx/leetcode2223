package ctx

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx,cancel := context.WithTimeout(ctx,time.Second)
	defer cancel()
	foo(ctx)
}

func foo(ctx context.Context) {
	var slowChan = make(chan struct{},1)
	go slow(slowChan)
	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
		return
	case <-slowChan:
		fmt.Println("slow done")
		return
	}
}

func slow(ch chan struct{}) {
	seed := rand.Intn(2000)
	fmt.Printf("seed = %v\n",seed)
	time.Sleep(time.Duration(time.Duration(seed) * time.Millisecond))
	ch <- struct{}{}
}