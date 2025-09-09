func mySqrt(x int) int {
    closest := 0
    for r := range x+1 {
        if (r)*(r) > x {break}
        closest = r
    }
    return closest
}

//Runtime: 18 ms, Beats 9.00%
//Memory: 4.27 MB, Beats 7.07%
