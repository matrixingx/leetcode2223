package spiralorder

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	fmt.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
	fmt.Println(spiralOrder([][]int{{6,9,7}}))
}