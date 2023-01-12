package threesumclosest

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4},1))
	fmt.Println(threeSumClosest([]int{0,0,0},1))
	fmt.Println(threeSumClosest([]int{4,0,5,-5,3,3,0,-4,-5},-2))
}