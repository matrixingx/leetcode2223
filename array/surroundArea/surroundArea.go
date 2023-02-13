package surroundarea

func solve(board [][]byte)  {
	var n,m = len(board),len(board[0])
	var vis [][]bool = make([][]bool, n)
	var dir = [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	for i := 0 ; i < n ; i ++ {
		vis[i] = make([]bool, m)
	}
	var markEdge func(n,m int)
	markEdge = func(n, m int) {
		for i := 0 ; i < m ; i ++ {
			if board[0][i] == 'O' {
				board[0][i] = 'A'
			}
			if board[n-1][i] == 'O' {
				board[n-1][i] = 'A'
			}
		}
		for i := 0 ; i < n ; i ++ {
			if board[i][0] == 'O' {
				board[i][0] = 'A'
			}
			if board[i][m-1] == 'O' {
				board[i][m-1] = 'A'
			}
		}
	}
	markEdge(n,m)
	var dfs func(i,j int) 
	dfs = func(i, j int) {
		for _,v := range dir {
			if (i+v[0]<0) || (i+v[0]>=n) ||(j+v[1]<0) || (j+v[1]>=m) {
				continue
			}
			newI,newJ := i+v[0],j+v[1]
			if board[newI][newJ] == 'O' {
				board[newI][newJ] = 'A'
				dfs(newI,newJ)
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'A' {
				dfs(i,j)
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'A' {
				board[i][j] = 'X'
			}
		}
	}
}