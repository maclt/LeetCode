func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    max := 0
    area := 0

    for left < right {
        area = min(height[left], height[right]) * (right - left)

        if area > max {
            max = area
        }

        if height[left] > height[right] {
            right = right -1 
        } else {
            left = left + 1
        }
    }

    return max
}
