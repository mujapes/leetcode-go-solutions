//expand is given a window of s containing a 2
// or 3 length palindrome and expands each end
// of the window until the ends are not equal
func expand(s string, longest *string, start, end int) {
    for {
        if end-start >= len(*longest) {*longest = s[start:end+1]}
        if start-1 < 0 || end+1 >= len(s) {break}
        start--
        end++
        if s[start] != s[end] {break}
    }
}

func longestPalindrome(s string) string {
    longest := s[:1]
    for i := range s {
        if i+1 < len(s) {
            if s[i+1] == s[i] {expand(s, &longest, i, i+1)}
            if i-1 >= 0 {
                if s[i-1] == s[i+1] {expand(s, &longest, i-1, i+1)}
            }
        }
    }
    return longest
}

//Runtime: 3 ms, Beats 61.29%
//Memory: 4.33 MB, Beats 89.13%
