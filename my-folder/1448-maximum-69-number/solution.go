func maximum69Number (num int) int {
    leftMostSix := -1
    cp := num
    for i := 0; cp > 0; i++ {
        if cp%10 == 6 {
            leftMostSix = i
        }
        
        cp /= 10
    }
    
    if leftMostSix == -1 {
        return num
    }
    
    return num + int(math.Pow(10.0, float64(leftMostSix)))*3
}
