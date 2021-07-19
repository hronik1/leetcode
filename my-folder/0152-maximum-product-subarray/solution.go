func maxProduct(nums []int) int {
    currentBestMax := 1
    currentBestMin := 1
    bestProduct := nums[0]
    for _, v := range nums {
        if v == 0 {
            currentBestMax = 1
            currentBestMin = 1
            if bestProduct < 0 {
                bestProduct = 0
            }
            continue
        }
        
        currentMax := currentBestMax*v
        currentMin := currentBestMin*v
        candidates := []int{currentMax, currentMin, v}
        currentBestMax = sliceMax(candidates)
        currentBestMin = sliceMin(candidates)
        if currentBestMax > bestProduct {
            bestProduct = currentBestMax
        }

    }
    
    return bestProduct
}

func sliceMax(nums []int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}

func sliceMin(nums []int) int {
    min := nums[0]
    for _, v := range nums {
        if v < min {
            min = v
        }
    }
    return min
}
