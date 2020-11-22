func countAndSay(n int) string {
    if n == 1 { 
        return "1"
    }
    
    return say(countAndSay(n-1))
}

func say(digits string) string {
    if len(digits) < 1 {
        return ""
    }
    
    var out strings.Builder
    count := 0
    prev := string(digits[0])
    for _, c := range digits {
        if string(c) == string(prev) {
            count++
        } else {
            out.WriteString(fmt.Sprintf("%d%s", count, prev))
            prev = string(c)
            count = 1
        }
    }
    
    out.WriteString(fmt.Sprintf("%d%s", count, prev))
    
    return out.String()
}
