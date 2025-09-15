import "strings"

func appendLetter(current string, depth int, digits *string, posLttrs *[][]string, cmbs *[]string, digToLttr *map[string]int) {
    if depth == len(*digits) {
        *cmbs = append(*cmbs, current)
        return
    }
    if (*digits)[depth:depth+1] == "7" || (*digits)[depth:depth+1] == "9" {
        for i := range 4 {appendLetter(current + (*posLttrs)[(*digToLttr)[(*digits)[depth:depth+1]]][i], depth+1, digits, posLttrs, cmbs, digToLttr)}
    } else {
        for i := range 3 {appendLetter(current + (*posLttrs)[(*digToLttr)[(*digits)[depth:depth+1]]][i], depth+1, digits, posLttrs, cmbs, digToLttr)}
    }
}

func letterCombinations(digits string) []string {
    if digits == "" {return []string{}}
    posLttrs := [][]string{
        {"a","b","c"},
        {"d","e","f"},
        {"g","h","i"},
        {"j","k","l"},
        {"m","n","o"},
        {"p","q","r","s"},
        {"t","u","v"},
        {"w","x","y","z"}}
    digToLttr := map[string]int{
        "2":0, 
        "3":1, 
        "4":2,
        "5":3,
        "6":4,
        "7":5,
        "8":6,
        "9":7}
    cmbsLen := 1
    for _ = range len(digits)-1 {cmbsLen *= 4}
    cmbs := make([]string, 0, cmbsLen)
    appendLetter("", 0, &digits, &posLttrs, &cmbs, &digToLttr)
    return cmbs
    
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.02 MB, Beats 55.69%
