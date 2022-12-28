package myatoi

import (
	"fmt"
	"strings"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	res := strings.TrimLeft("    000045  ", " ")
	fmt.Println(res)
	res = strings.TrimLeft(res, "0")
	fmt.Println(res)
	res2 := (2 << 30) - 1
	fmt.Println(res2)
	res3 := -(2 << 30)
	fmt.Println(res3)
	fmt.Println(myAtoi("    --91283472332.114514 with words"))
}
