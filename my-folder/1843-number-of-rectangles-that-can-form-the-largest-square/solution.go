func countGoodRectangles(rectangles [][]int) int {
    size := 0
    out := 0
    for _, dimensions := range rectangles {
        s := dimensions[0]
        if dimensions[1] < s {
            s = dimensions[1]
        }
        
        if s > size {
            size = s
            out = 0
        }
        
        if s == size {
            out++
        }
    }
    
    return out
}
