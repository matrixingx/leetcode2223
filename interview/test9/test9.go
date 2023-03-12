package test9

import (
	"strconv"
	"strings"
)

var mp = map[int]int {
	0:1,
	1:1024,
	2:1024*1024,
}

func solution(s string) int64 {
	maxValue := 2 << 32
	minValue := 0
	var integer int64
	var f float64
	var start = false
	var useFloat = false
	if strings.Contains(s,".") {
		useFloat = true
	}
	if len(strings.Split(s,".")) > 2 {
		return -1
	}
	var kmgCount = make([]int,3)
	var hasKMG = false
	var kmgIndex = 0
	for i := range s {
		if !start && (s[i] < '1' || s[i] > '9') {
			return -1
		} else if !start {
			start = true
			continue
		}
		if s[i] == '.' {
			continue
		}
		if !hasKMG {
			if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >='A' && s[i] <= 'Z')  {
				if s[i] != 'k' && s[i] != 'K' && s[i] != 'm' && s[i] != 'M' && s[i] != 'G' && s[i] != 'g' {
					return -1
				}
				switch s[i]{
				case 'k','K':
					kmgCount[0]++
				case 'm','M':
					kmgCount[1]++
				case 'g','G':
					kmgCount[2]++
				}
				hasKMG = true
				kmgIndex = i
				continue
			}
		}
		if hasKMG && (s[i] >= 'a' && s[i] <= 'z') || (s[i]>='A' && s[i] <= 'Z') {
			return -1
		}
	}
	if useFloat && kmgIndex != 0 {
		s1 := s[:kmgIndex]
		f,_ = strconv.ParseFloat(s1,10)
		for i := range kmgCount {
			if kmgCount[i] > 0 {
				temp := f*float64(mp[i])
				res := int(temp*1024)
				if res > maxValue || res < minValue {
					return -1
				}
				return int64(res)
			}
		}
	}
	if !useFloat && kmgIndex != 0 {
		s1 := s[:kmgIndex]
		integer,_ = strconv.ParseInt(s1,10,64)
		for i := range kmgCount {
			if kmgCount[i] > 0 {
				temp := integer*int64(mp[i])
				res := temp*1024
				if res > int64(maxValue) || res < int64(minValue) {
					return -1
				}
				return int64(res)
			}
		}
	}
	if kmgIndex == 0 {
		if useFloat {
			s1 := s
			f,_ = strconv.ParseFloat(s1,10)
			res := f
			if res > float64(maxValue) || res < float64(minValue) {
				return -1
			}
			return int64(res)
		} else {
			s1 := s
			integer,_ = strconv.ParseInt(s1,10,64)
			res := integer
			if res > int64(maxValue) || res < int64(minValue) {
				return -1
			}
			return int64(res)
		}
	}
	return -1
}