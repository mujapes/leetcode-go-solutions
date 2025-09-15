func threeSum(nums []int) [][]int {
    oneSums := make(map[int][]int)
    oneSums[nums[0]] = append(oneSums[nums[0]], 0)
    twoSums := make(map[[2]int]struct{})
    threeSums := make([][]int, 0, 100)
    var small, mid int
    for i, d := range nums {
        oneSums[d] = append(oneSums[d], i)
    }
    for i := 0; i < len(nums); i++ {
        if oneSums[nums[i]][0] != -1 {
            for j := i+1; j < len(nums); j++ {
                if oneSums[nums[j]][0] != -1 {
                    for _, d := range oneSums[-1*(nums[i]+nums[j])] {
                        if d == -1 {break}
                        if d != i && d != j {
                            switch min(nums[i], nums[j], -1*(nums[i]+nums[j])) {
                                case nums[i]:
                                    small = nums[i]
                                    if nums[j] <= -1*(nums[i]+nums[j]) {
                                        mid = nums[j]
                                    } else {mid = -1*(nums[i]+nums[j])}
                                case nums[j]:
                                    small = nums[j]
                                    if nums[i] <= -1*(nums[i]+nums[j]) {
                                        mid = nums[i]
                                    } else {mid = -1*(nums[i]+nums[j])}
                                default:
                                    small = -1*(nums[i]+nums[j])
                                    if nums[i] <= nums[j] {
                                        mid = nums[i]
                                    } else {mid = nums[j]}
                            }
                            if _, ok := twoSums[[2]int{small, mid}]; !ok {
                                twoSums[[2]int{small, mid}] = struct{}{}
                                        threeSums = append(threeSums, []int{small, mid, -1*(small+mid)})
                            }
                        }
                    }
                }
            }
            oneSums[nums[i]][0] = -1
        }
    }
    return threeSums
}

//Runtime: 690 ms, Beats 8.87%
//Memory: 9.74 MB, Beats 48.23%



/*func threeSum(nums []int) [][]int {
    oneSums := make(map[int][]int)
    oneSums[nums[0]] = append(oneSums[nums[0]], 0)
    twoSums := make(map[int][][3]int)
    threeSums := make([][3]int, 0, 100)
    var dup bool
    for i := 0; i < len(nums); i++ {
        if oneSums[nums[i]][0] != -1 {
            for j := i+1; j < len(nums); j++ {
                if i == 0 {
                    if _, ok := oneSums[nums[j]]; !ok {
                        oneSums[nums[j]] = append(oneSums[nums[j]], j)
                        for _, s := range twoSums[-1*nums[j]] {
                            if s[2] == 0 {
                                for _, thrSum := range threeSums {
                                    if thrSum[0] == nums[j] {
                                        if thrSum[1] == s[0] || thrSum[1] == s[1] || thrSum[2] == s[0] || thrSum[2] == s[1]
                                    }
                                }
                                s[2] = 1
                            }
                        }
                    }
                }
                dup = false
                for _, s := range twoSums[nums[i]+nums[j]] {
                    if nums[s[0]] == nums[i] || nums[s[1]] == nums[i] {
                        dup = true
                        break
                    }
                }
                if !dup {
                    twoSums[nums[i]+nums[j]] = append(twoSums[nums[i]+nums[j]], [2]int{i, j})
                    for _, idx := range oneSum[-1*(nums[i]+nums[j])] {
                        if idx != i && idx != j {
                            threeSums = append(threeSums, [3]int{nums[idx], nums[i], nums[j]})
                        }
                    }
                }
            }
            oneSums[nums[i]][0] = -1
        }
    }
    
}*/

/*func threeSum(nums []int) [][]int {
    var largeNums bool
    if len(nums) > 3 {largeNums = true}
    big1, big2, small1, small2 := -3000, -3000, 3000, 3000
    for _, n := range nums {
        if n > big1 {
            big2 = big1
            big1 = n
        } else if n > big2 {big2 = n}
        if n < small1 {
            small2 = small1
            small1 = n
        } else if n < small2 {small2 = n}
    }
    if big2 == -3000 {big2 = big1}
    if small2 == 3000 {small2 = small1}
    checkedDig := make(map[int]struct{})
    checkedSum := make(map[int]struct{})
    threeSums := make([][]int, 0, 20)
    var sum int
    for i := 0; i < len(nums); i++ {
        if _, ok := checkedDig[nums[i]]; !ok {
            if largeNums && (nums[i]+big1+big2 < 0 || nums[i]+small1+small2 > 0) {
                checkedDig[nums[i]] = struct{}{}
                continue
            }
            checkedSum = make(map[int]struct{})
            for _, s := range threeSums {
                for idx, n := range s {
                    if n == nums[i] {
                        if idx == 0 {
                            checkedSum[s[1]] = struct{}{}
                            checkedSum[s[2]] = struct{}{}
                        } else if idx == 1 {
                            checkedSum[s[0]] = struct{}{}
                            checkedSum[s[2]] = struct{}{}
                        } else {
                            checkedSum[s[0]] = struct{}{}
                            checkedSum[s[1]] = struct{}{}
                        }
                        break
                    }
                }
            }
            for j := i+1; j < len(nums); j++ {
                _, chk1 := checkedDig[nums[j]]
                _, chk2 := checkedSum[nums[j]]
                if !chk1 && !chk2 {
                    sum = nums[i]+nums[j]
                    for k := j+1; k < len(nums); k++ {
                        if sum + nums[k] == 0 {
                            checkedSum[nums[k]] = struct{}{}
                            threeSums = append(threeSums, []int{nums[i], nums[j], nums[k]})
                            break
                        }
                    }
                    checkedSum[nums[j]] = struct{}{}
                }
            }
            checkedDig[nums[i]] = struct{}{}
        }
    }
    return threeSums
}*/



/*import "slices"

type twoSum struct {
    idx1, idx2, val1, val2 int
}

func threeSum(nums []int) [][]int {
    twoSums := make(map[int][]twoSum)
    twoChecked := make(map[int]struct{})
    var sum int
    var collision bool
    for i := 0; i < len(nums); i++ {
        if _, ok := twoChecked[nums[i]]; !ok {
            for j := i+1; j < len(nums); j++ {
                sum = nums[i]+nums[j]
                collision = false
                for _, s := range twoSums[sum] {
                    if s.val1 == nums[i] || s.val2 == nums[i] {
                        collision = true
                        break
                    }
                }
                if !collision {twoSums[sum] = append(twoSums[sum], twoSum{i, j, nums[i], nums[j]})}
            }
            twoChecked[nums[i]] = struct{}{}
        }
    }
    threeChecked := make(map[int]struct{})
    threeSums := make([][]int, 0, 100)
    var twoSumCollision, threeSumCollision bool
    var curThreeSum []int
    var sml, mid, lrg int
    for i, n := range nums {
        if _, ok := threeChecked[n]; !ok {
            twoSumCollision = false
            for _, s := range twoSums[n*-1] {
                if s.idx1 == i || s.idx2 == i {
                    twoSumCollision = true
                    continue
                }
                if s.val1 < s.val2 {
                    if s.val2 < n {
                        sml, mid, lrg = s.val1, s.val2, n
                    } else if s.val1 < n {
                        sml, mid, lrg = s.val1, n, s.val2
                    } else {sml, mid, lrg = n, s.val1, s.val2}
                } else {
                    if s.val1 < n {
                        sml, mid, lrg = s.val2, s.val1, n
                    } else if s.val2 < n {
                        sml, mid, lrg = s.val2, n, s.val1
                    } else {sml, mid, lrg = n, s.val2, s.val1}
                }
                curThreeSum = []int{sml, mid, lrg}
                threeSumCollision = false
                for _, s := range threeSums {
                    if slices.Equal(curThreeSum, s) {
                        threeSumCollision = true
                        break
                    }
                }
                if !threeSumCollision {threeSums = append(threeSums, curThreeSum)}
            }

            if !twoSumCollision {threeChecked[n] = struct{}{}}
        }
    }
    return threeSums
}*/
