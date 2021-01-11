func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var concat1 strings.Builder
    var concat2 strings.Builder
    for _, s := range word1 {
        _, _ = concat1.WriteString(s)
    }
    for _, s := range word2 {
        _, _ = concat2.WriteString(s)
    }
    
    return concat1.String() == concat2.String()
}
