func moveZeroes(nums []int)  {
    var limit = len(nums)
    
    for i:=0; i<limit; i++ {
        if nums[i] == 0 {
            nums = append(nums[:i], nums[i+1:]...)
            nums = append(nums, 0)
            
            i = i -1
            limit = limit - 1
        }
    }
}
