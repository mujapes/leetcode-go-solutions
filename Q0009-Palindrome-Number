import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
    xConv := strconv.Itoa(x)
    xLen := len(xConv)
    if xLen == 1 {return true}
    for i := 0; i<xLen/2; i++ {
        if xConv[i] != xConv[xLen-1-i] {return false}
    }
    return true
}

//Runtime: 3ms, Beats 57.79%
//Memory: 6.24 MB, Beats 36.85%
