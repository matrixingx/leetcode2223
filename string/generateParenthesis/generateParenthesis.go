package generateParenthesis

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(ans string,leftIndex,rightIndex int)
	dfs = func (ans string,leftIndex,rightIndex int)  {
		if leftIndex == n && rightIndex == n {
			res = append(res, ans)
			return
		}
		if leftIndex < n {
			temp := ans + "("
			dfs(temp,leftIndex+1,rightIndex)
		}
		if rightIndex < leftIndex {
			temp := ans + ")"
			dfs(temp,leftIndex,rightIndex+1)
		}
	}
	dfs("",0,0)
	return res
}