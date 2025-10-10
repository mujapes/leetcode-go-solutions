func rotate(matrix [][]int)  {
    n := len(matrix)-1
    for layer := range (n+1)/2 {
        // 4 elements 90x degrees apart swapped clockwise, next 4 elements are
        // nearest clockwise neighbours, indent layer + reset rotation at >= 90 degrees
        for rot := layer; rot < n-layer; rot++ {
            matrix[layer][rot], matrix[rot][n-layer], matrix[n-layer][n-rot], matrix[n-rot][layer] = 
                matrix[n-rot][layer], matrix[layer][rot], matrix[rot][n-layer], matrix[n-layer][n-rot]
        }
    }
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.09 MB, Beats 97.35%
