import (
    "strings"
    "strconv"
)

func reverse(x int) int {
    var neg bool
    ovflw := []int{2,1,4,7,4,8,3,6,4,7}
    if x < 0 {
        neg = true
        x *= -1
    }
    xStr := strconv.Itoa(x)
    if len(xStr) == 10 {
        for i := range 10 {
            d, _ := strconv.Atoi(string(xStr[9-i]))
            if d > ovflw[i] {
                return 0
            } else if d < ovflw[i] {break}
        }
    }
    var xStrRev strings.Builder
    for i := len(xStr)-1; i >= 0; i-- {
        xStrRev.WriteByte(xStr[i])
    }
    if xRev, _ := strconv.Atoi(xStrRev.String()); neg {
        return -1 * xRev
    } else {return xRev}
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 3.99 MB, Beats 81.13%


/*func reverse(x int) int {
    if x == 0 {return 0}
    var neg bool
    if x < 0 {
        neg = true
        x *= -1
    }
    rev := make([]int, 0, 10)
    sum := 0
    ovflw := [5]int{21,47,48,36,48}
    var revDigs = [100]int{
	0, 10, 20, 30, 40, 50, 60, 70, 80, 90,
	1, 11, 21, 31, 41, 51, 61, 71, 81, 91,
	2, 12, 22, 32, 42, 52, 62, 72, 82, 92,
	3, 13, 23, 33, 43, 53, 63, 73, 83, 93,
	4, 14, 24, 34, 44, 54, 64, 74, 84, 94,
	5, 15, 25, 35, 45, 55, 65, 75, 85, 95,
	6, 16, 26, 36, 46, 56, 66, 76, 86, 96,
	7, 17, 27, 37, 47, 57, 67, 77, 87, 97,
	8, 18, 28, 38, 48, 58, 68, 78, 88, 98,
	9, 19, 29, 39, 49, 59, 69, 79, 89, 99,
    }
    var ovflwPoss bool
    var oddLen bool
    if x > 999999999 {ovflwPoss = true}
    dig := 0
    idx := 0
    for x > 0 {
        dig = x % 100
        if x >= 100 || dig >= 10 {
            dig = revDigs[dig]
        } else {oddLen = true}
        if ovflwPoss && dig > ovflw[idx] {return 0}
        if ovflwPoss && dig < ovflw[idx] {ovflwPoss = false}
        idx++
        rev = append(rev, dig)
        x /= 100
    }
    if ovflwPoss && !neg && dig == 48 {return 0}
    for i, d := range rev {
        if i == len(rev)-1 && oddLen {
            sum = 10*sum + d
        } else {sum = 100*sum + d}
    }
    if neg {sum *= -1}
    return sum
}*/

//Runtime: 0 ms, Beats 100.00%
//Memory: 3.93 MB, Beats 80.72%
