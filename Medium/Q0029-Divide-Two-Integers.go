func divide(dividend int, divisor int) int {
    if dividend == -2147483648 && divisor == -1 {return 2147483647}
    if divisor == 1 {return dividend}
    if divisor == -1 {return -dividend}
    quotient := 0
    if dividend < 0 {
        if divisor < 0 {
            for dividend <= divisor {
                dividend -= divisor
                quotient++
            }
        } else {
            for dividend <= -1 * divisor {
                dividend += divisor
                quotient--
            }
        }
    } else if divisor < 0 {
        for dividend >= -1 * divisor {
                dividend += divisor
                quotient--
            }
    } else {
        for dividend >= divisor {
                dividend -= divisor
                quotient++
            }
    }
    return quotient
}

// Runtime: 1260 ms, Beats 5.70%
// Memory: 4.32 MB, Beats 14.51%
