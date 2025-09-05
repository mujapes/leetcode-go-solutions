func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {return ""}
    commonPrefix := strs[0]
    if len(strs) > 1 {
        for _, str := range strs {
            if commonPrefix == "" {break}
            for i := 0; i<len(str); i++ {
                if i<len(commonPrefix) {
                    if str[i] == commonPrefix[i] {
                        continue
                    }
                }
                commonPrefix = str[:i]
                break
            }
            if len(str)<len(commonPrefix) {commonPrefix = str}
        }
    }
    return commonPrefix
}

//Runtime: 1 ms, Beats 9.98%
//Memory: 4.40 MB, Beats 31.88%
