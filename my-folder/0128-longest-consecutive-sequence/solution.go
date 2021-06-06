func longestConsecutive(nums []int) int {
    items := map[int]bool{}
    for _, v := range nums {
        items[v] = true
    }
    
    longest := 0
    for _, v := range nums {
        // only check subsequence length from first element in potential subsequence
        if ok := items[v-1]; ok {
            continue
        }
        
        if length := consecutiveLength(v, items); length > longest {
            longest = length
        }
    }
    
    return longest
}

func consecutiveLength(v int, items map[int]bool) int {
    length := 0
    for i := v; ; i++ {
        if ok := items[i]; !ok {
            break
        }
        
        length++
    }
    
    return length
}
