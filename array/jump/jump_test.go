package jump

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{1, 1, 2, 2, 0, 1, 1}))

}