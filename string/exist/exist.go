package exist

func exist(board [][]byte, word string) bool {
	var backTrack func(arr []byte,word string,i,j int,wordIdx int)
	var res = false
	var dir = [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	var n,m = len(board),len(board[0])
	var vis = make([][]bool,n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	backTrack = func(arr []byte, word string,i,j int,wordIdx int) {
		if res {
			return
		}
		if wordIdx == len(word) {
			res = true
			return
		}
		vis[i][j] = true
		for _,v := range dir {
			newI,newJ := i+v[0],j+v[1]
			if newI >= 0 && newI < n && newJ >= 0 && newJ < m &&
			 !vis[newI][newJ] && word[wordIdx] == board[newI][newJ]{
				arr = append(arr, board[newI][newJ])
				backTrack(arr,word,newI,newJ,wordIdx+1)
				arr = arr[:len(arr)-1]
			}
		}
		vis[i][j] = false
	}
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < m ; j++ {
			if !res && board[i][j] == word[0] {
				arr := []byte{board[i][j]}
				backTrack(arr,word,i,j,1)
			}
		}
	}
	return res
}