func getRow(rowIndex int) []int {
    var prevRow, curRow []int
    for r := range rowIndex+1 {
        if r > 0 {
            curRow = make([]int, 0, r+1)
            curRow = append(curRow, 1)
            for i := range prevRow {
                if i != 0 {curRow = append(curRow, prevRow[i-1]+prevRow[i])}
            }
            curRow = append(curRow, 1)
        } else {curRow = []int{1}}
        if r == rowIndex {return curRow}
        prevRow = curRow
    }
    return []int{}
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 3.90 MB, Beats 99.13%

/*
func getRow(rowIndex int) []int {
  if rowIndex == 0 {return []int{1}}
  curRow, prevRow := []int{1}, getRow(rowIndex-1)
  for i := range prevRow {
      if i == 0 {continue}
      curRow = append(curRow, prevRow[i-1]+prevRow[i])
  }
  return append(curRow, 1)
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.26 MB, Beats 7.56%
*/
