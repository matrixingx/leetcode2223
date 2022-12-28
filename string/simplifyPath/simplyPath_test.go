package simplifypath

import (
	"fmt"
	"testing"
)

func TestSimplyPath(t *testing.T) {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
