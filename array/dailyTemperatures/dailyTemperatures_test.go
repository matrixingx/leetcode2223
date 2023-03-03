package dailytemperatures

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	arr := []int{73,74,75,71,69,72,76,73,73,74}
	fmt.Println(dailyTemperatures(arr))
}