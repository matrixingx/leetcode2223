package maxarea

func maxArea(height []int) int {
	var res int
	var getArea func(i,j int) int
	var getMax func(i,j int) int 
	getArea = func(i, j int) int {
		if j < i {
			i,j = j,i
		}
		length := height[i]
		if height[j] < length {
			length = height[j]
		}
		return (j-i)*length
	}
	getMax =  func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(height) == 0 || len(height) == 1{
		return 0
	}
	var i,j = 0,len(height)-1
	for i != j {
		res = getMax(res,getArea(i,j))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}