var intToStr []string = []string{"0","1","2","3","4","5","6","7","8","9"}

func countAndSay(n int) string {
    if n == 1 {return intToStr[1]}
    encoding, prevDig, cnt, cntStr := "", rune(0), 0, ""
    for _, b := range countAndSay(n-1) {
        if prevDig == rune(0) {prevDig = b}
        if prevDig != b {
            cntStr = ""
            for cnt > 9 {
                cntStr = intToStr[cnt%10] + cntStr
                cnt /= 10
            }
            encoding += intToStr[cnt] + cntStr + string(prevDig)
            prevDig, cnt = b, 0
        }
        cnt++
    }
    cntStr = ""
    for cnt > 9 {
        cntStr = intToStr[cnt%10] + cntStr
        cnt /= 10
    }
    return encoding + intToStr[cnt] + cntStr + string(prevDig)
}

// Runtime: 8 ms, Beats 49.54%
// Memory: 9.22 MB, Beats 46.44%
