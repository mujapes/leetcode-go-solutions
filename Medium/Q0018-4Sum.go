func fourSum(nums []int, target int) [][]int {
    fourSums := make([][]int, 0, 100)
    checked := make(map[int]bool)
    twoSums := make(map[int][]int)
    var prevIdx int
    var dup, scnd bool
    for i := 0; i < len(nums); i++ {
        // Skip numbers that occur >=2 times before index i
        if chk, ok := checked[nums[i]]; !(ok && chk) {
            for j := i+1; j < len(nums); j++ {
                if chk, ok := checked[nums[j]]; !(ok && chk) {
                    dup, scnd = false, false
                    for cnt, idx := range twoSums[nums[i]+nums[j]] {
                        if idx == i || idx == j {dup = true}
                        if cnt%2 != 0 && dup {
                            if !scnd {
                                scnd = true
                                dup = false
                            } else {break}
                        }
                    }
                    if !dup {
                        for cnt, idx := range twoSums[target-(nums[i]+nums[j])] {
                            if idx == i || idx == j {
                                dup = true
                            }
                            if cnt%2 != 0{
                                if dup {
                                    dup = false
                                    continue
                                }
                                for _, s := range fourSums {
                                    if (s[0] == min(nums[i], nums[j]) && s[1] == max(nums[i], nums[j]) && s[2] == min(nums[idx], nums[prevIdx])) || (s[0] == min(nums[i], nums[idx]) && s[1] == max(nums[i], nums[idx]) && s[2] == min(nums[j], nums[prevIdx])) || (s[0] == min(nums[i], nums[prevIdx]) && s[1] == max(nums[i], nums[prevIdx]) && s[2] == min(nums[j], nums[idx])) || (s[0] == min(nums[j], nums[idx]) && s[1] == max(nums[j], nums[idx]) && s[2] == min(nums[i], nums[prevIdx])) || (s[0] == min(nums[j], nums[prevIdx]) && s[1] == max(nums[j], nums[prevIdx]) && s[2] == min(nums[i], nums[idx])) || (s[0] == min(nums[idx], nums[prevIdx]) && s[1] == max(nums[idx], nums[prevIdx]) && s[2] == min(nums[i], nums[j])) {
                                        dup = true
                                        break
                                    }
                                }
                                if !dup {fourSums = append(fourSums, []int{min(nums[i], nums[j]), max(nums[i], nums[j]), min(nums[idx], nums[prevIdx]), max(nums[idx], nums[prevIdx])})} else {dup = false}
                            } else {prevIdx = idx}
                        }
                        twoSums[nums[i]+nums[j]] = append(twoSums[nums[i]+nums[j]], i, j)
                    }
                }
            }
            // Set to false on 1st instance of nums[i] and true on 2nd
            checked[nums[i]] = ok
        }
    }
    return fourSums
}

//Runtime: 32 ms, Beats 17.83%
//Memory: 10.90 MB, Beats 5.62%
