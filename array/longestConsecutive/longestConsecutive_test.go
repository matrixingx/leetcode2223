package longestconsecutive

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
