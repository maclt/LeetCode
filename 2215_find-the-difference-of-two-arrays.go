import (
	"slices"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
    var new1 []int
    var new2 []int

    for i:=0; i<len(nums1); i++ {
        if !slices.Contains(nums2, nums1[i]) && !slices.Contains(new1, nums1[i]){
            new1 = append(new1, nums1[i])
        }
    }

    for i:=0; i<len(nums2); i++ {
        if !slices.Contains(nums1, nums2[i]) && !slices.Contains(new2, nums2[i]){
            new2 = append(new2, nums2[i])
        }
    }

    return [][]int{new1, new2}
}