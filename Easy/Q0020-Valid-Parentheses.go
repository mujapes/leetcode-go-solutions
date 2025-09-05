func isValid(s string) bool {
    openBrackets := make([]rune, 0, 40)
    complement := map[rune]rune{')':'(',']':'[','}':'{'}
    for _, b := range s {
        switch b {
            case '(', '[', '{':
                openBrackets = append(openBrackets, b)
            case ')', ']', '}':
                if len(openBrackets)>0 {
                    if openBrackets[len(openBrackets)-1] == complement[b] {
                        openBrackets = openBrackets[:len(openBrackets)-1]
                        continue
                    }
                }
                return false
        }
    }
    if len(openBrackets)!=0 {return false}
    return true
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.43 MB, Beats 38.47%
