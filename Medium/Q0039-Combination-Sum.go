func combinationSum(candidates []int, target int) [][]int {
    combs, combCpy := [][]int{}, []int{}
    var findCombs func(comb []int, start int, remaining int)
    findCombs = func(comb []int, start int, remaining int) {
        if remaining > 0 {
            for i := start; i < len(candidates); i++ {
                findCombs(append(comb, candidates[i]), i, remaining-candidates[i])
            }
        } else if remaining == 0 {
            combCpy = make([]int, len(comb))
            copy(combCpy, comb)
            combs = append(combs, combCpy)
        }
    }
    findCombs(make([]int, 0, 50), 0, target)
    return combs
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.86 MB, Beats 73.66%
