func findMaxConsecutiveOnes(nums []int) int {
    var maxLength int = 0;
    var length int = 0;

    for _, value := range nums {
        if value == 1 {
            length += 1
        } else {
            length = 0
        }
        
        if length > maxLength  {
            maxLength = length
        }
    }
    
    return maxLength
}
