import "math"

func reverse(x int) int {
    if x == 0 {return 0}
    var neg bool
    if x < 0 {
        neg = true
        x *= -1
    }
    rev := make([]int, 0, 10)
    sum := 0
    ovflw := []int{2,1,4,7,4,8,3,6,4,8}
    for x > 0 {
        rev = append(rev, x % 10)
        x /= 10
    }
    if len(rev) == 10 {
        for i, d := range rev {
            if d < ovflw[i] {break}
            if d > ovflw[i] {return 0}
            if i == 9 {
                if d == 8 {return 0}
            }
        }
    }
    for i, d := range rev {
        sum += int(math.Pow(10, float64(len(rev)-i-1))) * d
    }
    if neg {sum *= -1}
    return sum
}

//Runtime: 5 ms, Beats 16.88%
//Memory: 4.12 MB, Beats 4.82%
