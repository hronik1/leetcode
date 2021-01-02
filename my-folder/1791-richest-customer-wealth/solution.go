func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, banks := range accounts {
        curWealth := 0
        for _, balance := range banks {
            curWealth += balance
        }
        
        if curWealth > maxWealth {
            maxWealth = curWealth
        }
    }
    
    return maxWealth
}
