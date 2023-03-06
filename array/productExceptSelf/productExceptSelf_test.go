package productexceptself

import (
	"fmt"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}