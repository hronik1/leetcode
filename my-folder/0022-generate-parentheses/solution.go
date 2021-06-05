func generateParenthesis(n int) []string {
    return helper(n, "", 0, 0, []string{})
}

func helper(n int, s string, openCount int, closeCount int, res []string) []string {
    if openCount == n && closeCount == n {
        res = append(res, s)
        return res
    }
    
    if openCount < n {
        res = helper(n, fmt.Sprintf("%s(", s), openCount+1, closeCount, res)
    }
    
    if closeCount < openCount {
        res = helper(n, fmt.Sprintf("%s)", s), openCount, closeCount+1, res)
    }
    
    return res
}
