package removeduplicateletters

func removeDuplicateLetters(s string) string {
	var byteCount = make([]int,26)
	for i := range s {
		byteCount[s[i]-'a']++
	}
	var stack []byte
	var existSet = [26]bool{}
	for i := range s {
		v := s[i]
		if !existSet[v-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > v {
				last := stack[len(stack)-1] -'a'
				if byteCount[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				existSet[last] = false
			}
			stack = append(stack, v)
			existSet[v-'a'] = true
		}
		byteCount[v-'a']--
	}
	return string(stack)
}