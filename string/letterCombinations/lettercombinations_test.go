package lettercombinations

import (
	"fmt"
	"testing"
)

func TestLettercombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("2333"))
}