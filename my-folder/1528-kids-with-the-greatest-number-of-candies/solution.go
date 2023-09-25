func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := 1
    for _, candy := range candies {
        if candy > maxCandy {
            maxCandy = candy
        }
    }

    out := make([]bool, len(candies), len(candies))
    for i, candy := range candies {
        if candy + extraCandies >= maxCandy {
            out[i] = true
        }
    }

    return out
}
