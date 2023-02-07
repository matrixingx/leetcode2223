package generatetrees

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	res := generateTrees(3)
	fmt.Println(res)
}