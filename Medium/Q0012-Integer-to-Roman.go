import "strings"

func intToRoman(num int) string {
    numerals := []rune{'I','V','X','L','C','D','M'}
    var dig int
    pos := 6
    mod := 10000
    var roman strings.Builder
    for pos >= 0 {
        dig = num % mod
        mod /= 10
        dig /= mod
        for dig > 0 {
            switch {
                case dig == 9:
                    roman.WriteRune(numerals[pos])
                roman.WriteRune(numerals[pos+2])
                    dig -= 9
                case dig > 4:
                    roman.WriteRune(numerals[pos+1])
                    dig -= 5
                case dig == 4:
                    roman.WriteRune(numerals[pos])
                    roman.WriteRune(numerals[pos+1])
                    dig -= 4
                default:
                    roman.WriteRune(numerals[pos])
                    dig -= 1
            }
        }
        pos -= 2
    }
    return roman.String()
}

//Runtume: 0 ms, Beats 100.00%
//Memory: 5.01 MB, Beats 81.41%
