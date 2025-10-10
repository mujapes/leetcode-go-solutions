func jump(nums []int) int {
    if len(nums) == 1 {return 0}
    var cur, max2Jump, jumps int
    jumps++
    for cur+nums[cur] < len(nums)-1 {
        max2Jump = cur
        for i := cur; i <= cur+nums[cur]; i++ {
            if i+nums[i] > max2Jump+nums[max2Jump] {max2Jump = i}
        }
        jumps++
        cur = max2Jump
    }
    return jumps
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 7.86 MB, Beats 62.30%
