func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    
    out := make([]bool, len(candies))
    for i, v := range candies {
        if v + extraCandies >= max {
            out[i] = true
        }
    }
    
    return out
}
