func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums3 := make([]int, 0, len(nums1))
    var ptr1, ptr2 int
    for ptr1 != m || ptr2 != n {
        if ptr1 == m {
            nums3 = append(nums3, nums2[ptr2])
            ptr2++
        } else if ptr2 == n {
            nums3 = append(nums3, nums1[ptr1])
            ptr1++
        } else {
            if nums1[ptr1] < nums2[ptr2] {
                nums3 = append(nums3, nums1[ptr1])
                ptr1++
            } else {
                nums3 = append(nums3, nums2[ptr2])
                ptr2++
            }
        }
    }
    for i, n := range nums3 {
        nums1[i] = n
    }
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.24 MB, Beats 36.44%
