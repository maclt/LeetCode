func findNumbers(nums []int) int {
    var order int = 0
    var evenOrder int = 0
    
    for _, value := range nums {
        for value > 0 {
            value = value / 10
            order++
        }
        
        if order % 2 == 0 {
            evenOrder++
        }
        
        order = 0
    }
    
    return evenOrder
}