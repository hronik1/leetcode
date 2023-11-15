import "regexp"
import "strings"
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func isPalindrome(s string) bool {
    sanitized := sanitize(s)

    i := 0
    j := len(sanitized) - 1
    for i < j {
        if sanitized[i] != sanitized[j] {
            return false
        }

        i += 1
        j -= 1
    }

    return true
}

func sanitize(s string) string {
    return strings.ToLower(nonAlphanumericRegex.ReplaceAllString(s, ""))
}
