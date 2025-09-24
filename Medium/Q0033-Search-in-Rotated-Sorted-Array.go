func search(nums []int, target int) int {
    left, right, rotEnd, endVal := 0, len(nums)-1, -1, nums[len(nums)-1]
    if target == endVal {return len(nums)-1}
    if nums[0] > endVal {
        for left < right {
            mid := left+((right-left)/2)
            if nums[mid] == target {return mid}
            if nums[mid] > endVal {
                left = mid+1
                rotEnd = mid
            } else {right = mid}
        }
    }
    if target < endVal {
        left, right = max(0, rotEnd), len(nums)-1
    } else if rotEnd != -1 && target < nums[rotEnd] {
        left, right = 0, rotEnd
    }
    for left < right {
        mid := left+((right-left)/2)
        if nums[mid] == target {return mid}
        if nums[mid] < target {
            left = mid+1
        } else {right = mid}
    }
    return -1
}

// Runtime: 0 ms, Beats 100.00%
// Memory:4.49 MB, Beats 36.40%
