func findMaxAverage(nums []int, k int) float64 {
    var n int = len(nums)
    var ph int = 0
    var result int = 0

    for i:=0; i<k; i++ {
        ph = ph + nums[i]
    }

    result = ph

    for i:=1; i<n-k+1; i++ {
        var sum int = ph - nums[i-1] + nums[i+k-1]
        if sum > result {
            result = sum
        }
        ph = sum
    }

    return float64(result) / float64(k)
}
