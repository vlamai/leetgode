package _9_Combination_Sum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	r := dfs(0, 0, target, []int{}, candidates)
	res = append(res, r)
	return res
}

func dfs(i, total, target int, cur, candidates []int) []int {
	if total == target {
		return cur
	}
	if i >= len(candidates) || total > target {
		return nil
	}
	cur = append(cur, candidates[i])
	dfs(i, total+candidates[i], target, cur, candidates)
	dfs(i+1, total+candidates[i], target, cur[:len(cur)-1], candidates)
	return cur
}
