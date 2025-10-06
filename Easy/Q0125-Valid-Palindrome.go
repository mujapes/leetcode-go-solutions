import (
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    var stripped strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {stripped.WriteByte(b)}
    }
    s = stripped.String()
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    return true
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 5.10 MB, Beats 50.35%
