func maxArea(height []int) int {
    bestArea := 0
    for i, j := 0, len(height)-1; i < j; {
        lowest := height[i]
        if height[j] < height[i] {
            lowest = height[j]
        }
            
        area := lowest * (j-i)
        if area > bestArea {
            bestArea = area
        }
        
        if height[j] < height[i] {
            j--
        } else {
            i++
        }
    }   
    
    return bestArea
}
