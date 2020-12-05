func canPlaceFlowers(flowerbed []int, n int) bool {
    l := len(flowerbed)
    if l < n {
        return false
    }
    
    memo := map[int]int{}
    canPlaceFlowersHelper(flowerbed, l, 0, memo)
    
    return memo[0] >= n
}

func canPlaceFlowersHelper(flowerbed[]int, l int, i int, memo map[int]int) {
    if i >= l {
        memo[i] = 0
        return
    }
    
    if _, ok := memo[i]; ok {
        return
    }
    
    if flowerbed[i] == 1 {
        // a flower can't be placed in the next one, so that's why we inc by 2
        canPlaceFlowersHelper(flowerbed, l, i+2, memo)
        memo[i] = memo[i+2]
    } else {
        canPlaceFlowersHelper(flowerbed, l, i+1, memo)
        max := memo[i+1]
        if i == l-1 || flowerbed[i+1] != 1 {
            canPlaceFlowersHelper(flowerbed, l, i+2, memo)
            if max <= memo[i+2] {
                max = 1 + memo[i+2]
            }
        }
        memo[i] = max
    }
    
    return
}
