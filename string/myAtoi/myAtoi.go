package myatoi

import "strings"

func myAtoi(s string) int {
	s1 := strings.TrimLeft(s," ")
	var mark = true
	if len(s1) == 0 {
		return 0
	}
	if s1[0] == '-' {
		mark = false
	}
	var s2 = s1
	if s1[0] == '-' || s1[0] == '+' {
		s2 = s1[1:]
	}
	s2 = strings.TrimLeft(s2,"0")
	var res int64 = 0
	var max int64 = 2<<30-1
	var min int64 = 2<<30
	for i := 0 ; i < len(s2) ; i++ {
		if mark && res > max{
			return int(max)
		}
		if !mark && res > min {
			return int(-min)
		}
		if s2[i] < '0' || s2[i] > '9' {
			break
		}
		res *= 10
		res += int64(s2[i]-'0')
	}
	if mark && res > max{
		return int(max)
	}
	if !mark && res > min {
		return int(-min)
	}
	if mark {
		return int(res)
	} else {
		return -int(res)
	}
}