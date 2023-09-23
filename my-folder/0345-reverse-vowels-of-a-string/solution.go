
func reverseVowels(s string) string {
    i := 0
    sb := []rune(s)
    j := len(sb)-1
    for i < j {
        for i <len(sb) && !isVowel(s[i]) {
            i+=1
        }
        for j >= 0 && !isVowel(s[j]) {
            j-=1
        }
        if (i < j) {
            tmp := sb[i]
            sb[i] = sb[j]
            sb[j] = tmp
        }

        i+=1
        j-=1
    }

    return string(sb)
}

func isVowel(b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
} 
