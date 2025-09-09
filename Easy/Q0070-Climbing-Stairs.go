func climbStairs(n int) int {
    a, b := 1, 1
    for _ = range n-1 {a, b = b, a+b}
    return b
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 3.94 MB, Beats 35.37%
