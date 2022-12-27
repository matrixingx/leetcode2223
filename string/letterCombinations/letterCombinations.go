package lettercombinations

/*
	https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
*/

var mp = map[byte][]string{
	'2':{"a","b","c"},
	'3':{"d","e","f"},
	'4':{"g","h","i"},
	'5':{"j","k","l"},
	'6':{"m","n","o"},
	'7':{"p","q","r","s"},
	'8':{"t","u","v"},
	'9':{"w","x","y","z"},
}
// bfs做法
func letterCombinations(digits string) []string {
	var res []string
	for i := 0 ; i < len(digits) ; i++ {
		ans := mp[digits[i]]
		length := len(res)
		if length == 0 {
			for _,v := range ans {
				res = append(res, v)
			}
		} else {
			for j := 0 ; j < length ; j++ {
				temp := res[j]
				for _,v := range ans {
					temp2 := temp + v
					res = append(res, temp2)
				}
			}
			res = res[length:]
		}
	}
	return res
}