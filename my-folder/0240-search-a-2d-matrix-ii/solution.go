func searchMatrix(matrix [][]int, target int) bool {
    for _, row := range matrix {
        if binSearch(row, target) {
            return true
        }
    }
    
    return false
}

func binSearch(items []int, target int) bool {
    lo := 0
    hi := len(items) - 1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        
        if items[mid] == target {
            return true
        }
        
        if items[mid] > target {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    return false
}
