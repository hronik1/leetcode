func canPlaceFlowers(flowerbed []int, n int) bool {
    i := 0 
    for i < len(flowerbed) && n > 0 {
        if canPlace(flowerbed, i) {
            n -= 1
            i += 1
        }

        i += 1
    }

    return n == 0
}

func canPlace(flowerbed []int, i int) bool {
    // assume i is valid
    if flowerbed[i] == 1 {
        return false
    }

    if i-1 >= 0 && flowerbed[i-1] == 1 {
        return false
    }

    if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
        return false
    }

    return true
}
