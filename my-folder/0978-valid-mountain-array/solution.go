func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    
    if arr[1] < arr[0] {
        return false
    }
    
    increasing := true
    for i := 1; i < len(arr); i++ {
        if arr[i-1] == arr[i] {
            return false
        }
        
        if increasing {
            if arr[i-1] > arr[i] {
                increasing = false
            }
        } else {
            if arr[i-1] < arr[i] {
                return false
            }
        }
    }
    
    return !increasing
}
