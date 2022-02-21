func combinationSum(candidates []int, target int) [][]int {
	cur := make([]int, 0)
	res := make([][]int, 0)
	idx := 0
	dfs(candidates, target, cur, &res, idx)
	return res
}
func dfs(candidates []int, target int, cur []int, res *[][]int, idx int) {
	if target < 0 {
		return
	} else if target == 0 {
		//Method1:
		//tmp := make([]int, len(cur))
		//copy(tmp, cur)
		//*res = append(*res, tmp)

		//Method2:
		*res = append(*res, append([]int{}, cur...))
		//fmt.Printf("idx: %v, tmp: %v, res: %v \n", idx, tmp, res)
		return
	}
	for i := idx; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		dfs(candidates, target-candidates[i], cur, res, i)
		cur = cur[:len(cur)-1]
	}
	return
}