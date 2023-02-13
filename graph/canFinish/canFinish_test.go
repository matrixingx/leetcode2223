package canfinish

import (
	"fmt"
	"testing"
)

func TestCanfinish(t *testing.T) {
	numsCourse := 3
	prerequisites := [][]int{{1,0},{2,1},{2,0}}
	fmt.Println(canFinish(numsCourse,prerequisites))
	numsCourse2 := 3
	prerequisites2 := [][]int{{1,0},{2,1},{2,0},{0,1}}
	fmt.Println(canFinish(numsCourse2,prerequisites2))
}