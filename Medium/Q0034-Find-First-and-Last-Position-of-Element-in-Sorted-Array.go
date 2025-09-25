func searchRange(nums []int, target int) []int {
    left, right, startTarget, endTarget := 0, len(nums)-1, -1, -1
    if len(nums) > 0 && nums[len(nums)-1] == target {startTarget, endTarget = len(nums)-1, len(nums)-1}
    for left < right {
        mid := left+((right-left)/2)
        if nums[mid] == target {
           if endTarget == -1 {endTarget = mid}
            startTarget = mid
        }
        if nums[mid] < target {
            left = mid+1
        } else {right = mid}
    }
    if endTarget != -1 {
        if nums[len(nums)-1] == target {endTarget = len(nums)-1}
        left, right = endTarget, len(nums)-1
        for left < right {
            mid := left+((right-left)/2)
            if nums[mid] == target {
                endTarget = mid
            }
            if nums[mid] > target {
                right = mid
            } else {left = mid+1}
        }
    }
    return []int{startTarget, endTarget}
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 6.51 MB, Beats 5.05%
