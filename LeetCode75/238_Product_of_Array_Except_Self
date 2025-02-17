func productExceptSelf(nums []int) []int {
    product := 1
    var zeroPositions []int;

    for i:=0; i<len(nums); i++ {
        if nums[i] == 0 {
            zeroPositions = append(zeroPositions, i)
        } else {
            product = product * nums[i]
        }
    }

    answers := make([]int, len(nums))

    if len(zeroPositions) >= 2 {
        return answers
    }

    if len(zeroPositions) == 1 {
        answers[zeroPositions[0]] = product
        return answers
    }

    for i:=0; i<len(nums); i++ {
        answers[i] = product / nums[i]
    }

    return answers
}
