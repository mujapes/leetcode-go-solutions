import "strings"
/*
func convert(s string, numRows int) string {
    if numRows == 1 {return s}
    rows := make([][]rune, numRows)
    for i := range rows {
        rows[i] = make([]rune, 0, len(s)/numRows)
    }
    rowIdx := 0
    var fwd bool
    for _, c := range s {
        rows[rowIdx] = append(rows[rowIdx], c)
        if rowIdx == 0 {fwd = true}
        if rowIdx == numRows-1 {fwd = false}
        if fwd {
            rowIdx++
        } else {rowIdx--}
    }
    var zigZag strings.Builder
    for _, row := range rows {
        for _, r := range row {
            zigZag.WriteRune(r)
        }
    }
    return zigZag.String()
}*/

//Runtime: 0 ms, Beats 100.00%
//Memory: 8.43 MB, Beats 62.74%

func convert(s string, numRows int) string {
    if numRows == 1 {return s}
    var zigZag strings.Builder
    var idx int
    var fwd bool
    for row := range numRows {
        idx = row
        fwd = true
        for idx < len(s) {
            zigZag.WriteByte(s[idx])
            if row == 0 || row == numRows-1 {
                idx += 2*(numRows-1)
            } else if fwd {
                idx += 2*(numRows-row-1)
                fwd = false
            } else {
                idx += 2*(row)
                fwd = true
            }
        }
    }
    return zigZag.String()
}

//Runtime:0 ms, Beats 100.00%
//Memory: 5.89 MB, Beats 90.43%
