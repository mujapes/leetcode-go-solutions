func lengthOfLongestSubstring(s string) int {
    sRunes := []rune(s)
    substring := make([]rune, 1, 10)
    longestSubstring := 0
    count := 0
    duplicate := false
    for i, r := range sRunes {
        substring = []rune{r}
        count++
        for j := i+1; j<len(sRunes); j++ {
            for _, rSub := range substring {
                if rSub == sRunes[j] {duplicate = true}
            }
            if duplicate {break}
            substring = append(substring, sRunes[j])
            count++
        }
        if count > longestSubstring {
            longestSubstring = count
        }
        count = 0
        duplicate = false
    }
    return longestSubstring
}

//Runtine: 121 ms, Beats 13.80%
//Memory: 8.86 MB, Beats 11.99%
