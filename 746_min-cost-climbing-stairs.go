func minCostClimbingStairs(cost []int) int {
    length := len(cost)
    dp := make([]int, length)
    dp[0] = cost[0]
    dp[1] = cost[1]

    for i:=2; i<=length-1; i++ {
        dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
    }

    return min(dp[length-2], dp[length-1])
}
