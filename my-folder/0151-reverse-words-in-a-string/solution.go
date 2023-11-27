import "strings"

func reverseWords(s string) string {
    words := strings.Fields(s)
    for i := 0; i < len(words)/2; i += 1 {
        j := len(words) - 1 - i 
        tmp := words[i]
        words[i] = words[j]
        words[j] = tmp
    }

    return strings.Join(words, " ") 
}
