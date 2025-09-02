func twoSum(nums []int, target int) []int {
    diff := map[int]int{}
    for j, v := range nums {
        if i, ok := diff[v]; ok {
            return []int{i, j}
        }
        diff[target - v] = j
    }
    return []int{}
}

//Runtime: 1 ms, Beats 56.71%
//Memory: 6.05 MB, Beats 40.64%