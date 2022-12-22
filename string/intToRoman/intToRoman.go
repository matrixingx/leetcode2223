package inttoroman

import "strings"

func intToRoman(num int) string {
	var res = strings.Builder{}
	type roman struct {
		Value int
		Str string
	}
	var romans = []roman{
		{
			1000,"M",
		},
		{
			900,"CM",
		},
		{
			500,"D",
		},
		{
			400,"CD",
		},
		{
			100,"C",
		},
		{
			90,"XC",
		},
		{
			50,"L",
		},
		{
			40,"XL",
		},
		{
			10,"X",
		},
		{
			9,"IX",
		},
		{
			5,"V",
		},
		{
			4,"IV",
		},
		{
			1,"I",
		},
	}
	for _,v := range romans {
		for v.Value <= num {
			num -= v.Value
			res.WriteString(v.Str)
		}
		if num == 0 {
			break
		}
	}
	return res.String()
}