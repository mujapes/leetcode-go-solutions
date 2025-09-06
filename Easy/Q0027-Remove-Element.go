func removeElement(nums []int, val int) int {
    for i := 0; i<len(nums); i++ {
        if nums[i] == val {
            nums = append(nums[:i], nums[i+1:]...)
            i--
        }
    }
    return len(nums)
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.09 MB, Beats 70.24%
