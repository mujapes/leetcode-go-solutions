func maxArea(height []int) int {
    left := 0
    right := len(height)-1
    var mArea int
    var nArea int
    for left < right {
        if height[left] > height[right] {
            nArea = height[right]*(right-left)
            right--
        } else {
            nArea = height[left]*(right-left)
            left++
        }
        if nArea > mArea {mArea = nArea}
    }
    return mArea
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 9.45 MB, Beats 78.87%
