func reverseVowels(s string) string {
    var chars = []byte(s)
    var length int = len(chars)
    var found1 byte
    var found2 byte
    var position1 int = 0
    var position2 int = 0
    var vowels string = "aeiouAEIOU"

    for position1 + position2 <= length - 2 {
        if found1 != 0 {

        } else if strings.Contains(vowels, string(chars[position1])) {
            found1 = chars[position1]
        } else {
            position1 = position1 + 1
        }

        if found2 != 0 {

        } else if strings.Contains(vowels, string(chars[length - position2 - 1])) {
            found2 = chars[length - position2 - 1]
        } else {
            position2 = position2 + 1
        }

        if found1 != 0 && found2 != 0 {
            chars[position1], chars[length - position2 - 1] = found2, found1

            found1, found2 = 0, 0
            position1 = position1 + 1
            position2 = position2 + 1
        }
    }

    return string(chars[:])
}
