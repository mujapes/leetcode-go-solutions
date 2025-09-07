func searchInsert(nums []int, target int) int {
    if nums[0] >= target {return 0}
    if nums[len(nums)-1] == target {return len(nums)-1}
    if nums[len(nums)-1] < target {return len(nums)}
    start := 0
    end := len(nums)-1
    var mid int
    for end-start>1 {
        mid = start+((end-start)/2)
        if nums[mid] == target {return mid}
        if nums[mid] < target {
            start = mid
        } else {end = mid}
    }
    return end
}

//Runtime:0 ms, Beats 100.00%
//Memory: 4.89 MB, Beats 34.30% 
