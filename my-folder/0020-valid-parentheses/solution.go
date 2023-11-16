func isValid(s string) bool {
    seen := []rune{}
    for _, v := range s {
        if v == '(' || v == '{' || v == '[' {
            seen = append(seen, v)
        } else {
            if len(seen) == 0 {
                return false
            }
            
            lastI := len(seen) - 1
            popped := seen[lastI]
            seen = seen[:lastI]
            if !matchingPair(popped, v) {
                return false
            }
        }
    }

    if len(seen) != 0 {
        return false
    }

    return true
}

func matchingPair(opening rune, closing rune) bool {
    if opening == '(' {
        return closing == ')'
    } else if opening == '{' {
        return closing == '}'
    } else {
        return closing == ']'
    }
}
