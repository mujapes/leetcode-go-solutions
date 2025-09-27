func generate(numRows int) [][]int {
    pascalTriangle := make([][]int, 0)
    for r := range numRows {
        curRow := make([]int, 0, r+1)
        if r > 0 {
            curRow = append(curRow, 1)
            for i, n := range pascalTriangle[r-1] {
                if i != r-1 {
                    curRow = append(curRow, n+pascalTriangle[r-1][i+1])
                }
            }
            curRow = append(curRow, 1)
        } else {curRow = append(curRow, 1)}
        pascalTriangle = append(pascalTriangle, curRow)
    }
    return pascalTriangle
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.38 MB, Beats 48.94%
