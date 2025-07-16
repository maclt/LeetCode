func isSubsequence(s string, t string) bool {
    index := 0
    if len(s) == 0 {
        return true
    }

    if len(t) == 0 {
        return false
    }
    
    for i:=0; i<len(t); i++ {
        if t[i] == s[index] {
            index = index + 1
        }

        if len(s) == index {
            return true
        }
    }

    return false
}
