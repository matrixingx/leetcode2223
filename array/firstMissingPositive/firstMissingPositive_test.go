package firstmissingpositive

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{-1,4,2,1,9,10}))

	fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
	fmt.Println(firstMissingPositive([]int{100,101,103,1,3,4,2,5,7}))

}