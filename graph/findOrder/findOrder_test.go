package findorder

import (
	"fmt"
	"testing"
)

func TestFindOrder(t *testing.T) {
	numsCourse := 4
	prerequisites := [][]int{{1,0},{2,0},{3,1},{3,2}}
	fmt.Println(findOrder(numsCourse,prerequisites))
}