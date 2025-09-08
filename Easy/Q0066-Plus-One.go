func plusOne(digits []int) []int {
    i := len(digits)-1
    for digits[i] == 9 {
        digits[i] = 0;
        if i == 0 {
            digits = append([]int{0}, digits...)
        } else {i--}
    }
    digits[i]++
    return digits 
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 3.91 MB, Beats 78.28%
