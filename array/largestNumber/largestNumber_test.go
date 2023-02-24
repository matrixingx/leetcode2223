package largestnumber

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	nums := []int{2, 10}
	nums2 := []int{3, 30, 34, 5, 9}
	nums3 := []int{113, 1130, 1133, 1134}
	fmt.Println(largestNumber(nums))
	fmt.Println(largestNumber(nums2))
	fmt.Println(largestNumber(nums3))

}
