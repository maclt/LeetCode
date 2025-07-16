func pivotIndex(nums []int) int {
    var left = 0
    var right = sumWithForLoop(nums) - nums[0]

    for i:=0; i<len(nums); i++ {
        if left == right {
            return i
        }

        if i == len(nums) -1 {
            break
        }

        left = left + nums[i]
        right = right - nums[i+1]
    }

    return -1
}

func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}
