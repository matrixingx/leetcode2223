package minwindow

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC","ABC"))
	fmt.Println(minWindow("a","a"))
	fmt.Println(minWindow("a","aa"))
}