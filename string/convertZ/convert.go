package convert

import "strings"

func convert(s string, numRows int) string {
	var res = strings.Builder{}
	var length = len(s)
	if numRows == 1 || length <= numRows {
		return s
	}
	var t = 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j = j + t {
			_ = res.WriteByte(s[j+i])
			if i > 0 && i < numRows - 1 && j+t-i < length {
				_ = res.WriteByte(s[j+t-i])
			}
		}
	}
	return res.String()
}