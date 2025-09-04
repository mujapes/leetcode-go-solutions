func romanToInt(s string) int {
    sum := 0
    sLen := len(s)
    for i := 0; i<sLen; i++ {
        switch s[i] {
            case 'M':
                sum += 1000
            case 'D':
                sum += 500
            case 'C':
                if i<=sLen-2 {
                    switch s[i+1] {
                        case 'M':
                            sum += 900
                            i++
                        case 'D':
                            sum += 400
                            i++
                        default:
                            sum += 100
                    }
                } else {sum += 100}
            case 'L':
                sum += 50
            case 'X':
                if i<=sLen-2 {
                    switch s[i+1] {
                        case 'C':
                            sum += 90
                            i++
                        case 'L':
                            sum += 40
                            i++
                        default:
                            sum += 10
                    }
                } else {sum += 10}
            case 'V':
                sum += 5
            case 'I':
                if i<=sLen-2 {
                    switch s[i+1] {
                        case 'X':
                            sum += 9
                            i++
                        case 'V':
                            sum += 4
                            i++
                        default:
                            sum += 1
                    }
                } else {sum += 1}
        }
    }
    return sum
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.82 MB, Beats 18.87%
