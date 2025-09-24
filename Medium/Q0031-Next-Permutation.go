func nextPermutation(nums []int)  {
    for i := len(nums)-1; i >= 0; i-- {
        if i > 0 && (i == 1 || nums[i-1] < nums[i]) {
            left := i-1
            right := len(nums)-1
            for nums[left] >= nums[right] && left < right {right--}
            if left >= right {right = len(nums)-1}
            for left < right {
                nums[left], nums[right] = nums[right], nums[left]
                if nums[left] <= nums[right] {right--} else {right = len(nums)-1}
                left++
            }
            break
        }
    }

    // Runtime: 0 ms, Beats 100.00%
    // Memory: 4.32 MB, Beats 37.50%
    
    /*
    [1,2,3], 
    [1,3,2], 
    [2,1,3], 
    [2,3,1], 
    [3,1,2], 
    [3,2,1]

    1234
    1243
    1324
    1342
    1423
    1432
    2134
    2143
    2314
    2341
    2413
    2431
    3124
    3142
    3214
    3241
    3412
    3421
    4123
    4132
    4213
    4231
    4312
    4321
    */
}
