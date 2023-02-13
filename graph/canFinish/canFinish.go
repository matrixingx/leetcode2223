package canfinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edge = make([][]int,numCourses)
		visited = make([]int,numCourses)
		valid = true
		dfs func(i int)
	)
	for _,v := range prerequisites {
		edge[v[1]] = append(edge[v[1]], v[0])
	}
	dfs = func(i int) {
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
	}
	for i := 0 ; i < numCourses ; i++ {
		dfs(i)
		if !valid {
			return valid
		}
	}
	return valid
}