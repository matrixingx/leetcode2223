package numislands

import (
	"fmt"
	"testing"
)

var grid1 = [][]byte{
	{'1', '1', '1', '1', '0'},
	{'1', '1', '0', '1', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '0', '0', '0'},
}

var grid2 = [][]byte{
	{'1', '1', '0', '0', '0'},
	{'1', '1', '0', '0', '0'},
	{'0', '0', '1', '0', '0'},
	{'0', '0', '0', '1', '1'},
}

func TestNumIslands(t *testing.T) {
	fmt.Println(numIslands(grid1))
	fmt.Println(numIslands(grid2))
}
