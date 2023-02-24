package largestnumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if nums[0] == 0 {
        return "0"
    }
	var numsString []string
	for i := range nums {
		numsString = append(numsString, strconv.Itoa(nums[i]))
	}
	sort.SliceStable(numsString,func(i, j int) bool {
		a,b := numsString[i]+numsString[j],numsString[j]+numsString[i]
		aI,_ := strconv.Atoi(a)
		bI,_ := strconv.Atoi(b)
		return aI>bI
	})
	var res strings.Builder
	for i := range numsString {
		res.WriteString(numsString[i])
	}
	if res.String()[0] == '0' {
		return "0"
	}
	return res.String()
}