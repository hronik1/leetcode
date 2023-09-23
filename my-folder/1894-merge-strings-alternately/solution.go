import "strings"
func mergeAlternately(word1 string, word2 string) string {
    var sb strings.Builder
    i := 0
    j := 0
    index := 0
    for i < len(word1) && j < len(word2) {
        if index%2 == 0 {
            _ = sb.WriteByte(word1[i])
            i+=1
        } else {
            _ = sb.WriteByte(word2[j])
            j+=1 
        }

        index += 1
    }

    for i < len(word1) {
        _ = sb.WriteByte(word1[i])
        i+=1
    }

    for j < len(word2) {
        _ = sb.WriteByte(word2[j])
        j+=1 
    }

    return sb.String()
}
