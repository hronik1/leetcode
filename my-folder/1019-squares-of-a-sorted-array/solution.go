func sortedSquares(nums []int) []int {
    //find smallest non-negative index
    nonNegI := 0
    for nonNegI < len(nums) && nums[nonNegI] < 0 {
        nonNegI += 1
    }

    negI := nonNegI - 1
    var squares []int
    for negI >= 0 && nonNegI < len(nums) {
        abs := -nums[negI]
        nonNeg := nums[nonNegI]
        if abs > nonNeg {
            squares = append(squares, nonNeg*nonNeg)
            nonNegI += 1
        } else {
            squares = append(squares, abs*abs)
            negI -= 1
        }
    }
    
    for nonNegI < len(nums) {
        v := nums[nonNegI]
        squares = append(squares, v*v)
        nonNegI += 1
    }    
    
    for negI >= 0 {
        v := nums[negI]
        squares = append(squares, v*v)
        negI -= 1
    }  
    
    return squares
}
