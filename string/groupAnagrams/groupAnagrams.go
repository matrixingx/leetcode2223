package groupanagrams

func groupAnagrams(strs []string) [][]string {
	var mp = make(map[[26]int64][]string)
	var res [][]string
	for i := range strs {
		k := [26]int64{}
		for _,v := range strs[i] {
			k[v-'a']++
		}
		mp[k] = append(mp[k], strs[i])
	}
	for _,v := range mp {
		res = append(res, v)
	}
	return res
}