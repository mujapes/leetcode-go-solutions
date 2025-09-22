func generateParenthesis(n int) []string {
    if n == 1 {return []string{"()"}}
    parentheses := make([]string, 0)
    added := make(map[string]struct{})
    for _, p := range generateParenthesis(n-1) {
        parentheses = append(parentheses, "(" + p + ")")
    }
    for m := range n/2 {
        par1 := generateParenthesis(m+1)
        par2 := generateParenthesis(n-m-1)
        for _, p1 := range par1 {
            for _, p2 := range par2 {
                if _, ok := added[p1 + p2]; !ok {
                    parentheses = append(parentheses, p1 + p2)
                    added[p1 + p2] = struct{}{}
                }
                if _, ok := added[p2 + p1]; !ok {
                    parentheses = append(parentheses, p2 + p1)
                    added[p2 + p1] = struct{}{}
                }
            }
        }
    }
    return parentheses
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 5.37 MB, Beats 13.87%
