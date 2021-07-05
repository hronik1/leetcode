func backspaceCompare(s string, t string) bool {
    sStack := buildRuneStack(s)
    tStack := buildRuneStack(t)
    
    if len(sStack) != len(tStack) {
        return false
    }
    
    for i, r := range sStack {
        if r != rune(tStack[i]) {
            return false
        }
    }
    
    return true
}

func buildRuneStack(s string) []rune {
    stack := []rune{}
    for _, r := range s {
        if r == '#' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, r)
        }
    }
    
    return stack
}
