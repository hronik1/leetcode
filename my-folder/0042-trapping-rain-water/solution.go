func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    
    maxHeightLeft := make([]int, len(height))
    maxHeightLeft[0] = height[0]
    for i := 1; i < len(height); i++ {
        maxHeightLeft[i] = maxHeightLeft[i-1]
        if height[i] > maxHeightLeft[i] {
            maxHeightLeft[i] = height[i]
        }
    }
    
    maxHeightRight := make([]int, len(height))
    maxHeightRight[len(height)-1] = height[len(height)-1]
    for i := len(height)-2; i >= 0; i-- {
        maxHeightRight[i] = maxHeightRight[i+1]
        if height[i] > maxHeightRight[i] {
            maxHeightRight[i] = height[i]
        }
    }
    
    vol := 0
    for i := 0; i < len(height)-1; i++ {
        minHeight := maxHeightLeft[i]
        if minHeight > maxHeightRight[i] {
            minHeight = maxHeightRight[i]
        }
        
        if minHeight > height[i] {
            vol += (minHeight-height[i])
        }
    }
    
    return vol
}
