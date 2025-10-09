import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
    res := make([][]int, 0, len(candidates))
    var findCombs func(comb []int, start int, remaining int)
    combCpy := []int{}
	findCombs = func(comb []int, start int, remaining int) {
        if remaining == 0 {
            combCpy = make([]int, len(comb))
			copy(combCpy, comb)
			res = append(res, combCpy)
        } else if remaining > 0 {
            for i := start; i < len(candidates); i++ {
                dupCnt := 1
                for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
                    dupCnt++
                    i++
                }
                for copies := range dupCnt {
                    findCombs(append(comb, candidates[i+1-dupCnt:i+2+copies-dupCnt]...), i+1, remaining-(candidates[i]*(copies+1)))
                }
            }
        }
    }
    findCombs(make([]int, 0, len(candidates)), 0, target)
    return res
}

// Runtime: 2 ms, Beats 30.22%
// Memory: 4.42 MB, Beats 87.72%
