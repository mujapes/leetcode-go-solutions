func pow(x, e int) int {
    result := 1
    for _ = range e {
        result *= x
    }
    return result
}

func myAtoi(s string) int {
    digits := "123456789"
    sInts := make([]int, 0, len(s))
    var start bool
    var neg bool
    var sign bool
    var nonDig bool
    for pos, d := range s {
        if d == '0' {
            if start {
                sInts = append(sInts, 0)
            }
        } else if d == '-' {
            if !start && !sign {
                if pos > 0 {if s[pos-1] == '0' {break}}
                neg = true
                sign = true
            } else {break}
        } else if d == '+' {
            if !start && !sign {
                if pos > 0 {if s[pos-1] == '0' {break}}
                sign = true
            } else {break}
        } else if d == ' ' {
            if !start && !sign {
                if pos > 0 {if s[pos-1] == '0' {break}}
            } else {break}
        } else {
            for i, dig := range digits {
                nonDig = true
                if d == dig {
                    sInts = append(sInts, i+1)
                    if !start {start = true}
                    nonDig = false
                    break
                }
            }
            if nonDig {break}
        }
    }
    var ovflw []int
    if neg {
        ovflw = []int{2,1,4,7,4,8,3,6,4,8}
    } else {ovflw = []int{2,1,4,7,4,8,3,6,4,7}}
    if len(sInts) == 10 {
        for i, d := range sInts {
            if d > ovflw[i] {
                if neg {return -2147483648}
                return 2147483647
            } else if d < ovflw[i] {break}
        }
    } else if len(sInts) > 10 {
        if neg {return -2147483648}
        return 2147483647
    }
    sum := 0
    for i, d := range sInts {
        sum += pow(10, len(sInts) - i - 1) * d
    }
    if neg {sum *= -1}
    return sum
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.35 MB, Beats 10.60%
