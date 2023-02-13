package surroundarea

import (
	"fmt"
	"testing"
)

var matrix = [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}

func TestSolve(t *testing.T) {
	solve(matrix)
	fmt.Println(matrix)
}