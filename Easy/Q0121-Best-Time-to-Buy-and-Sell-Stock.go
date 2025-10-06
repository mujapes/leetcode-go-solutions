func maxProfit(prices []int) int {
    left, right, minLeft, maxRight, maxProf := 0, len(prices)-1, 0, len(prices)-1, 0
    if prices[left] < prices[right] {maxProf = prices[right]-prices[left]}
    for left < right {
        if prices[maxRight]-prices[left+1] > prices[right-1]-prices[minLeft] {
            left++
            if prices[left] < prices[minLeft] {minLeft = left}
            if prices[maxRight]-prices[left] > maxProf {maxProf = prices[maxRight]-prices[left]}
        } else {
            right--
            if prices[right] > prices[maxRight] {maxRight = right}
            if prices[right]-prices[minLeft] > maxProf {maxProf = prices[right]-prices[minLeft]}
        }
    }
    return maxProf
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 9.68 MB, Beats 79.27%
