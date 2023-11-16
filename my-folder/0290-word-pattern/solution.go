import "strings"
func wordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
    if len(pattern) != len(words) {
        return false
    }

    wordToRune := map[string]byte{}
    runeToWord := map[byte]string{}
    for i, word := range words {
        r := pattern[i]
        if mappedWord, ok := wordToRune[word]; ok {
            if mappedWord != r {
                return false
            }
        } else if _, ok := runeToWord[r]; ok {
            return false
        }

        wordToRune[word] = r
        runeToWord[r] = word
    } 

    return true
}
