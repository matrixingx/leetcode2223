package minwindow

func minWindow(s string, t string) string {
	var ori = make(map[byte]int)
	for i := range t {
		ori[t[i]]++
	}
	var cnt = make(map[byte]int)
	var ansL,ansR = -1,-1
	var length = 1<<32
	for left,right := 0,0 ; right < len(s) ; right ++ {
		cnt[s[right]]++
		for check(ori,cnt) && left <= right {
			if right - left +1 < length {
				length = right - left + 1
				ansL,ansR = left,left+length
			}
			cnt[s[left]]--
			left++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func check(ori,cnt map[byte]int) bool {
	for k,v := range ori {
		if cntV,ok := cnt[k];ok {
			if v > cntV {
				return false
			}
		} else {
			return false
		}
	}
	return true
}