import (
    "slices"
)

func gcdOfStrings(str1 string, str2 string) string {
    var dividers1 = getDeviders(str1)
    var dividers2 = getDeviders(str2)

    var cd = intersection(dividers1, dividers2)
    slices.Reverse(cd)

    for _, n := range cd {
        if str1[0:n] == str2[0:n] {
            return str1[0:n]
        }
    }

    return ""
}

func getDeviders(str string) []int {
    var n = len(str)
    var deviders []int

    for i:=1; i <= n; i++ {
        if n%i == 0 {
            var substr string
            var flag bool = true

            for j:=1; j<=n/i; j++ {
                if j!=1 {
                    if substr != str[i*(j-1) : i*j]{
                        flag = false
                        break
                    }
                }
                substr = str[i*(j-1) : i*j]
            }

            if flag {
                deviders = append(deviders, i)
            }
        }
    }

    return deviders
}

func intersection(nums1 []int, nums2 []int) []int {
    var result []int

    set := make(map[int]struct{})
    for _, v := range nums2 {
        set[v] = struct{}{}
    }

    for _, v := range nums1 {
        if _, ok := set[v]; ok {
            result = append(result, v)
        }
    }

    return result
}