func myPow(x float64, n int) float64 {
    if n == 0 || x == 1.0 {return 1.0}
    if x == 0.0 {return 0.0}
    var neg bool
    if n < 0 {
        neg = true
        n *= -1
    }
    var log2n int
    res := x
    for doublings := range 31 {
        if n >= 2 << doublings {
            res *= res
            log2n++
        }
    }
    for _ = range n-(1<<log2n) {res *= x}
    if neg {res = 1/res}
    return res
}

// Runtime: 1910 ms, Beats 5.07%
// Memory: 4.02 MB, Beats 26.04%
