func fizzBuzz(n int) []string {
    out := []string{}
    for i := 1; i <= n; i++ {
        cur := ""
        if i%3 == 0 {
            cur += "Fizz"
        }
        if i%5 == 0 {
            cur += "Buzz"
        }
        if cur == "" {
            cur = fmt.Sprintf("%d", i)
        }
        
        out = append(out, cur)
    }
    
    return out
}
