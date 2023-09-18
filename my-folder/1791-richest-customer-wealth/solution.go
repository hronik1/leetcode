func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, banks := range accounts {
        wealth := 0
        for _, v := range banks {
            wealth = wealth + v
        }

        if wealth > maxWealth {
            maxWealth = wealth
        } 
    }

    return maxWealth
}
