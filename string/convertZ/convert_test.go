package convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	res := convert("PAYPALISHIRING",3)
	fmt.Println(res)
}