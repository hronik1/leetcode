import "sort"

func maxCoins(piles []int) int {
    sort.Slice(piles, func(i, j int) bool {
        return piles[i] > piles[j]
    })
    
    sum := 0
    stoppingPoint := len(piles)*2/3
    for i := 0; i < stoppingPoint; i++ {
        if i%2 == 1 {
            sum += piles[i]
        }
    }
    
    return sum
}
