func lengthOfLastWord(s string) int {
    count := 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            if count > 0 {return count}
        } else {count++}
    }
    return count
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.09 MB, Beats 83.48%
