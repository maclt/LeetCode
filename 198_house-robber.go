func rob(nums []int) int {
    var n int = len(nums)

    if n == 1 {
        return nums[0]
    }

    maxSum := make([]int, n)
    maxSum[0] = nums[0]
    maxSum[1] = max(nums[0], nums[1])

    for i:=2; i<n; i++ {
        maxSum[i] = max(maxSum[i-1], maxSum[i-2]+nums[i])
    }

    return maxSum[n-1]
}
