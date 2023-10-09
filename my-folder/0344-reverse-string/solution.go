func reverseString(s []byte) {
    midI := len(s)/2
    for i := 0; i < midI; i += 1 {
        tmp := s[i]
        s[i] = s[len(s)-1-i] 
        s[len(s)-1-i] = tmp
    }
}
