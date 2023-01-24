package merge

import "sort"

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.SliceStable(intervals,func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) == 0 {
		return res
	}
	res = append(res, intervals[0])
	var n = len(intervals)
	for i := 1 ; i < n ; i ++ {
		v := res[len(res)-1]
		if v[1] >= intervals[i][0] {
			res[len(res)-1] = []int{v[0],max(v[1],intervals[i][1])}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}