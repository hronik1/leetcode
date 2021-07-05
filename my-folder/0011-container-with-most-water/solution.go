func maxArea(height []int) int {
    maxArea := 0
    lo, hi := 0, len(height)-1
    for lo < hi {
        l := hi-lo
        h := height[lo]
        if height[hi] < h {
            h = height[hi]
            hi--
        } else {
            lo++
        }
        
        area := h * (l)
        if area > maxArea {
            maxArea = area
        }
    }
    
    return maxArea
}
