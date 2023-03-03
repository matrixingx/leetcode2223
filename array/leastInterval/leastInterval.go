package leastinterval

func leastInterval(tasks []byte, n int) int {
	var taskLen = len(tasks)
	var taskMap = make(map[byte]int)
	for i := range tasks {
		taskMap[tasks[i]]++
	}
	var maxColumn,maxTaskCount = 0,0
	for _,v := range taskMap {
		if v > maxColumn {
			maxColumn = v
			maxTaskCount = 1
		} else if v == maxColumn {
			maxTaskCount++
		}
	}
	res := (maxColumn-1)*(n+1)+maxTaskCount
	if res > taskLen {
		return res
	}
	return taskLen
}