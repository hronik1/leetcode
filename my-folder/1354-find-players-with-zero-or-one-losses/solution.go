func findWinners(matches [][]int) [][]int {
    onlyWins := map[int]bool{}
    singleLoss := map[int]bool{}
    multipleLosses := map[int]bool{}
    for _, match := range matches {
        winner := match[0]
        loser := match[1]
        if !singleLoss[winner] && !multipleLosses[winner] {
            onlyWins[winner] = true
        }

        delete(onlyWins, loser)
        if !singleLoss[loser] && !multipleLosses[loser] {
            singleLoss[loser] = true
        } else {
            delete(singleLoss, loser)
            multipleLosses[loser] = true
        }
    }

    winners := []int{}
    for k, _ := range onlyWins {
        winners = append(winners, k)
    }

    singleLosers := []int{}
    for k, _ := range singleLoss {
        singleLosers = append(singleLosers, k)
    }

    slices.Sort(winners)
    slices.Sort(singleLosers)

    out := [][]int{winners, singleLosers}
    return out
}
