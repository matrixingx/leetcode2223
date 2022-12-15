package connect1

import (
	"fmt"
	"testing"
)

var testNode = &Node{
	Val:1,
	Left: &Node{
		Val:2,
		Left: &Node{
			Val:4,
		},
		Right: &Node{
			Val:5,
		},
	},
	Right: &Node{
		Val:3,
		Left: &Node{
			Val:6,
		},
		Right: &Node{
			Val:7,
		},
	},
}

func TestConnect(t *testing.T) {
	res := connect(testNode)
	if res != nil {
		fmt.Println(res)
	}
}