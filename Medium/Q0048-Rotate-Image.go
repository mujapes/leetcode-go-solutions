func rotate(matrix [][]int)  {
    topLeft, topRight, botRight, botLeft := []int{0, 0}, []int{0, len(matrix)-1}, []int{len(matrix)-1, len(matrix)-1}, []int{len(matrix)-1, 0}
    for topLeft[1] < topRight[1] {
        // 4 pointers have values swapped (clockwise) then move clockwise.
        // Once topLeft pointer reaches original topRight position break 
        // and set pointers to corners of inner matrix
        for topLeft[1] < topRight[1] {
            matrix[topLeft[0]][topLeft[1]], matrix[topRight[0]][topRight[1]],  matrix[botRight[0]][botRight[1]], matrix[botLeft[0]][botLeft[1]] = matrix[botLeft[0]][botLeft[1]], matrix[topLeft[0]][topLeft[1]], matrix[topRight[0]][topRight[1]],  matrix[botRight[0]][botRight[1]]
            
            topLeft[1]++
            topRight[0]++
            botRight[1]--
            botLeft[0]--
        }
        topLeft[0], topLeft[1], topRight[0], topRight[1], botRight[0], botRight[1], botLeft[0], botLeft[1] = botLeft[0]+1, botLeft[1]+1, topLeft[0]+1, topLeft[1]-1, topRight[0]-1, topRight[1]-1, botRight[0]-1, botRight[1]+1
        // If inner matrix 1x1 or non-existant topLeft and topRight pointers 
        // will be in the same spot or at previous botRight and botLeft positions
    }
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.12 MB, Beats 60.41%
