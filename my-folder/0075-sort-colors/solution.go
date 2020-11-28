func sortColors(nums []int)  {
    numCount := []int{0,0,0}
    for _, v := range nums {
        // would error check here if out of range
        numCount[v]++
    }
    
    index := 0
    for color, count := range numCount {
        for i := 0; i < count; i, index = i+1, index+1 {
            nums[index] = color 
        }
    }
}
