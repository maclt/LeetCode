func longestOnes(nums []int, k int) int {
    i := 0;
    j := 0;
    max := 0;

    for j<len(nums) {
        if nums[j] == 0 {
            k--;
        }

        for k<0 {
            if nums[i] == 0 {
                k++;
            }
            i++;
        }

        if max < j-i+1 {
            max = j-i+1;
        }

        j++;
    }

    return max;
}