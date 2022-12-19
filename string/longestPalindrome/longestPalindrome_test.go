package longestpalindrome

import (
	"testing"
	"fmt"
)

func TestLongestpalindrome(t *testing.T) {
	s := "abcdefg"
	fmt.Println(s[0:3])
	fmt.Println(s[3:6])
	res := longestPalindrome("ab")
	fmt.Println(res)
}