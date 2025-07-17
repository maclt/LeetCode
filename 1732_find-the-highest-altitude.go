func largestAltitude(gain []int) int {
    var highest int = 0
    var height int = 0
    var n int = len(gain)

    for i:=0; i<n; i++ {
        height = height + gain[i]
        if highest < height {
            highest = height
        }
    }

    return highest
}