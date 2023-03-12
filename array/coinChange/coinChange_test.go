package coinchange

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{2,5,10,1},27))
	fmt.Println(coinChange([]int{2},3))
	fmt.Println(coinChange([]int{1},0))

}