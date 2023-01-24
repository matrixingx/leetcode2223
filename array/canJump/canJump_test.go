package canjump

import (
	"fmt"
	"testing"
)

func TestCanjump(t *testing.T) {
	fmt.Println(canJump([]int{1, 1, 2, 2, 0, 1, 1}))

	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{3, 2, 0, 0, 5}))
}
