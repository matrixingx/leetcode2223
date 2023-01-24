package combinationsum2

import (
	"fmt"
	"testing"
)

func TestConbinationSum(t *testing.T) {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5},8))
	fmt.Println(combinationSum2([]int{2,5,2,1,2},5))
}