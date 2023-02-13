package findorder

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edge = make([][]int,numCourses)
		visited = make([]int,numCourses)
		res []int
		valid = true
		dfs func(i int)
	)
	for _,v := range prerequisites {
		edge[v[1]] = append(edge[v[1]], v[0])
	}
	dfs = func (i int)  {
		visited[i] = 1
		for _,v := range edge[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[i] = 2
		res = append(res, i)
	}
	for i := 0 ; i < numCourses ; i ++ {
		if visited[i] == 0 {
			dfs(i)
		}
		if !valid {
			return []int{}
		}
	}
	for i := 0 ; i < numCourses/2 ; i++ {
		res[i],res[numCourses-i-1] = res[numCourses-i-1],res[i]
	}
	return res
}