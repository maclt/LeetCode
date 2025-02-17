func mergeAlternately(word1 string, word2 string) string {
    var length1 int = len(word1)
    var length2 int = len(word2)

    var minLength = min(length1, length2)
    
    var output string

    for i := 0; i< minLength; i++ {
        output = output + string(word1[i]) + string(word2[i])
    }

    if minLength == length1 && minLength == length2 {
    } else if minLength == length1 {
        output = output + word2[minLength:]
    } else if minLength == length2 {
        output = output + word1[minLength:]
    }

    return output
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
