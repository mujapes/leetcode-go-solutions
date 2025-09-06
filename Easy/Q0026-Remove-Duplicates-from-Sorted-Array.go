func removeDuplicates(nums []int) int {
    var count int
    for i := 0; i<len(nums); i++ {
        count = 0
        for j := i+1; j<len(nums); j++ {
            if nums[i] == nums[j] {
                count++
            } else {break}
        }
        if count>0 {
            nums = append(nums[:i], nums[i+count:]...)
        }
    }
    return len(nums)
}
//Runtime: 0 ms, Beats 100.00%
//Memory: 6.18 MB, Beats 95%
