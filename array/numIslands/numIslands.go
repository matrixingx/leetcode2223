package numislands

func numIslands(grid [][]byte) int {
	var n,m = len(grid),len(grid[0])
	var visited = make([][]bool,n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	var dir = [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	var dfs func(i,j int,start bool)
	var res int
	dfs = func(i, j int,start bool) {
		if visited[i][j] {
			return
		} else {
			visited[i][j] = true
		}
		if grid[i][j] == '0' {
			return
		}
		for _,v := range dir {
			if (i+v[0]<0) || (i+v[0]>=n) ||(j+v[1]<0) || (j+v[1]>=m) {
				continue
			}
			newI,newJ := i+v[0],j+v[1]
			dfs(newI,newJ,false)
		}
		if start {
			res++
		}
	}
	for i := range grid {
		for j := range grid[i] {
			dfs(i,j,true)
		}
	}
	return res
}