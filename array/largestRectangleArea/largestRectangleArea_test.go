package largestrectanglearea

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{1,1}))
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}