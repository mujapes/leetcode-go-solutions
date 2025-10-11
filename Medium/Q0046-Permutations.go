func permute(nums []int) [][]int {
    var res [][]int
    var heapPermutation func(size int)
    // Implement heap's algorithm
    heapPermutation = func(size int) {
        if size == 1 {
            res = append(res, make([]int, len(nums)))
            copy(res[len(res)-1], nums)
            return
        }
        for i := range size-1 {
            heapPermutation(size-1)
            if size & 1 == 1 {
                nums[0], nums[size-1] = nums[size-1], nums[0]
            } else {
                nums[i], nums[size-1] = nums[size-1], nums[i]
            }
        }
        heapPermutation(size-1)
    }
    heapPermutation(len(nums))
    return res
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.69 MB, Beats 55.48%
