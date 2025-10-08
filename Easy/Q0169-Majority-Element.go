func majorityElement(nums []int) int {
    res, freq := nums[0], 0
    for _, num := range nums {
        if num != res {
            freq--
            if freq < 0 {
                res = num
                freq = 1
            }
        } else {freq++}
    }
    return res
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 8.76 MB, Beats 30.27%
