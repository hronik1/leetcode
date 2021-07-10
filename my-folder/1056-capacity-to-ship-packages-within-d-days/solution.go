func shipWithinDays(weights []int, days int) int {
    hi := 0
    lo := 0
    for _, v := range weights {
        if v > lo {
            lo = v
        }
        
        hi += v
    }
    
    lowestCapacity := hi
    for lo <= hi {
        mid := lo + (hi-lo)/2
        
        if isValidCapacity(weights, days, mid) {
            if mid < lowestCapacity {
                lowestCapacity = mid
            }
            hi = mid -1
        } else {
            lo = mid+1
        }
    }
    
    return lowestCapacity
}

func isValidCapacity(weights []int, days int, capacity int) bool {
    day := 0
    load := 0
    for _, weight := range weights {
        if load+weight > capacity {
            day++
            load = 0
        }
        
        if day >= days {
            return false
        }
        
        load+=weight 
    }
    return true
}
