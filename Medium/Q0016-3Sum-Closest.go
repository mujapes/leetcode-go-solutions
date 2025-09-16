import "slices"

func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
    closest := nums[0]+nums[1]+nums[2]
    if closest == target {return target}
    var distance int
    if closest < target {
        distance = target-closest
    } else {distance = closest-target}
    checked := make(map[int]struct{})
    iPrev := nums[0]+1
    var sum int
    var j, k, jPrev, kPrev, possClosest int
    for i := 0; i < len(nums)-2; i++ {
        if nums[i] == iPrev {continue}
        j, k = i+1, len(nums)-1
        jPrev, kPrev = nums[j]+1, nums[k]+1
        for {
            for _, ok := checked[nums[j]]; ok || nums[j] == jPrev; {
                jPrev = nums[j]
                j++
                if j >= k {break}
            }
            for _, ok := checked[nums[k]]; ok || nums[k] == kPrev; {
                kPrev = nums[k]
                k--
                if j >= k {break}
            }
            if j >= k {break}
            possClosest = nums[i]+nums[j]+nums[k]
            sum += possClosest
            if possClosest == target {
                return target
            } else if possClosest > target {
                if possClosest-target < distance {
                    distance = possClosest-target
                    closest = possClosest
                }
                kPrev = nums[k]
                k--
            } else {
                if target-possClosest < distance {
                    distance = target-possClosest
                    closest = possClosest
                }
                jPrev = nums[j]
                j++
            }
        }
        checked[nums[i]] = struct{}{}
        iPrev = nums[i]
    }
    return closest
}

//Runtime: 80 ms, Beats 10.42%
//Memory: 5.75 MB, Beats 4.20%
