func strStr(haystack string, needle string) int {
    for i := len(needle) - 1; i < len(haystack); i += 1 {
        //fmt.Printf("i:%d, l(hay):%d", i, len(haystack))
        if haystack[i] == needle[len(needle)-1] {
            start := i + 1 - len(needle)
            if haystack[start:i+1] == needle {
                return start
            }
        }
    }

    return -1
}
