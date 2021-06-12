
func addBinary(a string, b string) string {
    maxLen := len(a)
    if len(b) > maxLen {
        maxLen = len(b)
    }
    
    paddedA := fmt.Sprintf("%0*s", maxLen, a)
    paddedB := fmt.Sprintf("%0*s", maxLen, b)
    
    out := strings.Builder{}
    carry := 0
    for i := maxLen-1; i >= 0; i-- {
        cur := carry + RToI(rune(paddedA[i])) + RToI(rune(paddedB[i]))
        carry = cur/2
        cur = cur%2
        out.WriteRune(IToR(cur))
    }
    
    if carry > 0 {
        out.WriteRune(IToR(carry))
    }
    
    return Reverse(out.String())
}

func RToI(r rune) int {
    return int(r - '0')
}

func IToR(i int) rune {
    if i == 0 {
        return '0'
    } else {
        return '1'
    }
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

