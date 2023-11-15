import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    s := strconv.Itoa(x)
    i := 0
    j := len(s) - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }

        i += 1
        j -= 1
    }

    return true

}
