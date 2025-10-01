func singleNumber(nums []int) int {
    mask := 0
    for _, n := range nums {mask = mask ^ n}
    return mask
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 8.01 MB, Beats 54.91%
