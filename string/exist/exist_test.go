package exist

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	var board = [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	fmt.Println(exist(board,"ABCB"))
}