func convert(s string, numRows int) string {
    if numRows == 1 || len(s) <= numRows {
        return s
    }

	zigZag := make(map[int]string)
    row := 0
    isDownward := true

    for _, char := range s {
        zigZag[row] = zigZag[row] + string(char)

        if isDownward {
            if row == numRows-1 {
                isDownward = false
                row--
            } else {
                row++
            }
        } else {
            if row == 0 {
                isDownward = true
                row++
            } else {
                row--
            }
        }
    }

    var ans string

    for i:=0; i<numRows; i++ {
        ans = ans + zigZag[i]
    }

    return ans
}
