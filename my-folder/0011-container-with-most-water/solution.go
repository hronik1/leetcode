func maxArea(height []int) int {
    lo := 0
    hi := len(height) - 1
    area := 0
    for lo < hi {
        curHeight := height[lo]
        if height[hi] < curHeight {
            curHeight = height[hi]
        }
        
        curArea := (hi - lo) * curHeight
        if curArea > area {
            area = curArea
        }

        if height[lo] < height[hi] {
            lo += 1
        } else {
            hi -= 1
        }
    }

    return area
}
