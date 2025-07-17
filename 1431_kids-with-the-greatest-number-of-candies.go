func kidsWithCandies(candies []int, extraCandies int) []bool {
    var max = slices.Max(candies)
    var result [] bool = make([]bool, len(candies))
    for i:=0; i<len(candies); i++ {
        if candies[i] + extraCandies >= max{
            result[i] = true
        } else {
            result[i] = false
        }
    }

    return result
}