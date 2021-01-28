func halvesAreAlike(s string) bool {
    v := map[rune]bool{
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
        'A': true,
        'E': true,
        'I': true,
        'O': true,
        'U': true,
    }
    h := len(s)/2
    firstHalfCount := 0
    secondHalfCount := 0
    for i, r := range s {
        if isVowel := v[r]; isVowel {
            if i < h {
                firstHalfCount++
            } else {
                secondHalfCount++
            }
        }
    }
    
    return firstHalfCount == secondHalfCount
}
