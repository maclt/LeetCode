func findKthLargest(nums []int, k int) int {
    var m = make(map[int]int);

    for i:=0; i<len(nums); i++ {
        val := nums[i]
        m[val] = m[val] + 1;
    }

    counter := 0;

    for val:=10000; val>=-10000; val-- {
        counter = counter + m[val];

        if counter >= k {
            return val;
        }
    }

    return -10000;
}
