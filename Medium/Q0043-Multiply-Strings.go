func multiply(num1 string, num2 string) string {
    runeToInt := map[rune]int{'0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9}
    intToString := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
    prod, digProd, res := make([]int, len(num1)+len(num2)), 0, ""
    var addToProd func(val, offset int)
    addToProd = func(val, offset int) {
        overflow := 1
        if val > 10 {
            prod[offset-1] += val/10
            for prod[offset-overflow] > 9 {
                prod[offset-overflow-1] += prod[offset-overflow] / 10
                prod[offset-overflow] = prod[offset-overflow]%10
                overflow++
            }
            val = val%10
        }
        overflow = 0
        prod[offset] += val
        for prod[offset-overflow] > 9 {
            prod[offset-overflow-1] += prod[offset-overflow] / 10
            prod[offset-overflow] = prod[offset-overflow]%10
            overflow++
        }
    }
    for i, d1 := range num1 {
        for j, d2 := range num2 {
            digProd = runeToInt[d1]*runeToInt[d2]
            addToProd(digProd, i+j+1)
        }
    }
    start := false
    for _, d := range prod {if d != 0 || start{res += intToString[d];start = true}}
    if res == "" {return "0"}
    return res
}

// Runtime: 8 ms, Beats 17.22%
// Memory: 5.12 MB, Beats 23.45%
