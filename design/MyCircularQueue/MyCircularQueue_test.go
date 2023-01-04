package mycircularqueue

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue(t *testing.T) {
	queue := Constructor(3)
	fmt.Println(queue.EnQueue(1))
	fmt.Println(queue.EnQueue(2))
	fmt.Println(queue.EnQueue(3))
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.Rear())
	fmt.Println(queue.IsFull())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.Rear())
}