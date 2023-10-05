func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sToT := map[byte]byte{}
    tToS := map[byte]byte{}
    for i := 0; i < len(s); i += 1 {
        sByte := s[i]
        tByte := t[i]
        if val, ok := sToT[sByte]; ok {
            if val != tByte {
                return false
            }
        } else if _, ok := tToS[tByte]; ok {
            return false
        }
        
        sToT[sByte] = tByte
        tToS[tByte] = sByte
    }
    
    return true
}
