import "math"
func increasingTriplet(nums []int) bool {
    var min int = math.MaxInt32;
    var secondMin int = math.MaxInt32;

    for i:=0; i< len(nums); i++ {
        var num = nums[i];
        
        if num < min {
            min = num;
            continue;
        }

        if num > min && num < secondMin {
            secondMin = num;
            continue;
        }

        if num > secondMin {
            return true;
        }
    }

    return false;
}
