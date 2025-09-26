func isValidSudoku(board [][]byte) bool {
    var rowSeen, colSeen, boxSeen [100]bool
    for row, r := range board {
        for col, b := range r {
            if int(b) > 47 && int(b) < 58 {
                v := int(b)-48
                if !rowSeen[(row*10)+v] {
                    rowSeen[(row*10)+v] = true
                } else {return false}
                if !colSeen[(col*10)+v] {
                    colSeen[(col*10)+v] = true
                } else {return false}
                if !boxSeen[(3*int(row/3)+col/3)*10+v] {
                    boxSeen[(3*int(row/3)+col/3)*10+v] = true
                } else {return false}
            } 
        }
    }
    return true
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.64 MB, Beats 79.75%
