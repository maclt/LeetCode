func removeStars(s string) string {
    i := 0
    stack := make([]byte, 0)

    for i<len(s) {
        if s[i] == '*' {
            stack = stack[:len(stack)-1];
        } else {
            stack = append(stack, s[i])
        }

        i++;
    }

    return string(stack);
}