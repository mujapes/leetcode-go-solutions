func convertToTitle(columnNumber int) string {
    intToLetter := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    title := ""
    for columnNumber > 0 {
        columnNumber--
        title = intToLetter[columnNumber%26] + title
        columnNumber /= 26
    }
    return title
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 3.95 MB, Beats 20.45%
