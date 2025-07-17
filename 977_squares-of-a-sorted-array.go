func sortedSquares(nums []int) []int {
    squaredNums := make([]int, len(nums))

    for i, value := range nums {
        squaredNums[i] = value * value
    }
    
    sort.Ints(squaredNums)
    
    return squaredNums
}