func groupAnagrams(strs []string) [][]string {
    var res [][]string
    var sorted string
    var added bool
    for _, str := range strs {
        sorted = ""
        added = false
    iterateStr:
        for _, strCh := range str {
            for i, srtCh := range sorted {
                if strCh < srtCh {
                    sorted = sorted[:i] + string(strCh) + sorted[i:]
                    continue iterateStr
                }
            }
            sorted += string(strCh)
        }
        for i := range res {
            if res[i][0] == sorted {
                res[i] = append(res[i], str)
                added = true
                break
            }
        }
        if !added {res = append(res, []string{sorted, str})}
    }
    for i := range res {res[i] = res[i][1:]}
    return res
}

// Runtime: 94 ms, Beats 5.07%
// Memory: 9.83 MB, Beats 46.85%
