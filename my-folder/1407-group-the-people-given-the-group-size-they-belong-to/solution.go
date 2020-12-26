func groupThePeople(groupSizes []int) [][]int {
    out := [][]int{}
    sizeMapping := map[int]int{} 
    currPosition := 0
    
    for person, size := range groupSizes {
        var outI int
        var ok bool
        if outI, ok = sizeMapping[size]; !ok {
            out = append(out, []int{})
            outI = currPosition
            sizeMapping[size] = outI
            currPosition++
        }
        
        out[outI] = append(out[outI], person)
        if len(out[outI]) == size {
            delete(sizeMapping, size)
        }
    }
    
    return out
}
