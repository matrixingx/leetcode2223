package simplifypath

import "strings"

/*
	https://leetcode.cn/problems/simplify-path/
*/

func simplifyPath(path string) string {
	var stack []string
	paths := strings.Split(path,"/")
	for i := range paths {
		if paths[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		} else if paths[i] == "" || paths[i] == "." {
			continue
		} else {
			stack = append(stack, paths[i])
		}
	}
	var res = strings.Builder{}
	if len(stack) == 0 {
		return "/"
	}
	for i := range stack {
		res.WriteString("/")
		res.WriteString(stack[i])
	}
	return res.String()
}