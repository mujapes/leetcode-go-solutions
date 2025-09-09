import "strings"

func addBinary(a string, b string) string {
    var j int
    sum := make([]int, len(a))
    for i, r := range a {
        if r == '1' {
            sum[i] = 1
        } else {sum[i] = 0}
    }
    if len(b) > len(a) {
        padding := make([]int, len(b)-len(a))
        sum = append(padding, sum...)
    }
    for i := 1; i <= len(b); i++ {
        if b[len(b)-i] == '1' {
            j = len(sum)-i
            for sum[j] == 1 {
                sum[j] = 0
                if j == 0 {
                    sum = append([]int{0}, sum...)
                } else {j--}
            }
            sum[j] = 1
        }
    }
    
    var sumConv strings.Builder
    for _, v := range sum {
        if v == 0 {
            sumConv.WriteRune('0')
        } else {sumConv.WriteRune('1')}
    }
    return sumConv.String()
}

//Runtime: 0 ms, Beats 100;00%
//Memory: 4.33 MB, Beats 66.22%
