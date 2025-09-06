func strStr(haystack string, needle string) int {
    var same bool
    for i := 0; i<len(haystack); i++ {
        if len(needle) <= len(haystack)-i {
            for j := 0; j<len(needle); j++ {
                same = true
                if haystack[i+j] != needle[j] {
                    same = false
                    break
                }
            }
            if same {return i}
        }
    }
    return -1
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 3.92 MB, Beats 48.85%
