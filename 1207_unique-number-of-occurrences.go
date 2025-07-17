import (
	"slices"
)

func uniqueOccurrences(arr []int) bool {
    occurences := make(map[int]int)

    for i:=0; i<len(arr); i++ {
        occurences[arr[i]] = occurences[arr[i]] + 1
    }

    var values []int
    for _, v := range occurences {
        if slices.Contains(values, v) {
            return false
        }
        values = append(values, v)
    }

    return true
}