import "sort"

func deckRevealedIncreasing(deck []int) []int {
    sort.Slice(deck, func(i int, j int) bool {
        return deck[i] < deck[j]
    })
    
    index := make([]int, len(deck))
    for i := 0; i < len(deck); i++ {
        index[i] = i
    }
    
    out := make([]int, len(deck))
    for i := 0; i < len(deck); i++ {
        out[index[0]] = deck[i]
        index = index[1:]
        if len(index) > 0 {
            first := index[0]
            index = index[1:]
            index = append(index, first)
        }
    }

    return out
}
