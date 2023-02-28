package solution

func solution2(arr []int64) []int64 {
	var bitMap = make([]byte,1<<31)
	for i := range arr {
		if bitMap[arr[i]] <=3 {
			bitMap[arr[i]]++
		}
	}
	var res []int64
	for i := range bitMap {
		if bitMap[i] == 2 {
			res = append(res, int64(i))
		}
	}
	return res
}