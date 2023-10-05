func findRestaurant(list1 []string, list2 []string) []string {
    indices := map[string]int{}
    for i, v := range list1 {
        indices[v] = i
    }
    
    curBestSum := -1
    res := []string{}
    for j, v := range list2 {
        if i, ok := indices[v]; ok {
            curSum := i + j
            if len(res) == 0 || curBestSum == curSum {
                res = append(res, v)
                curBestSum = curSum
            } else if curSum < curBestSum {
                res = []string{v}
                curBestSum = curSum
            }
        }
    }
    
    return res
}
