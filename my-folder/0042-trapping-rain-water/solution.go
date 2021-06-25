func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    
    maxLeft := make([]int, len(height))
    maxLeft[0] = height[0]
    for i := 1; i < len(height); i++ {
        max := height[i]
        if maxLeft[i-1] > max {
            max = maxLeft[i-1]
        }
        
        maxLeft[i] = max
    }
    
    maxRight := make([]int, len(height))
    maxRight[len(height)-1] = height[len(height)-1]
    for i := len(height)-2; i >= 0; i-- {
        max := height[i]
        if maxRight[i+1] > max {
            max = maxRight[i+1]
        }
        
        maxRight[i] = max
    }
    
    out := 0
    for i := 0; i < len(height); i++ {
        minMax := maxLeft[i]
        if maxRight[i] < minMax {
            minMax = maxRight[i]
        }
        
        out += (minMax-height[i])
    }
    
    return out
    
}
