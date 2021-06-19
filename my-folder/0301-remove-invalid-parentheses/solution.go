func removeInvalidParentheses(s string) []string {
    length := requiredLength(s)
    solutions := map[string]bool{}
    chars := make([]rune, length)
    
    helper(s, 0, 0, 0, length, chars, 0, solutions)
    
    out := make([]string, 0, len(solutions))
    for k, _ := range solutions {
        out = append(out, k)
    }
    
    return out
}

func helper(s string, i int, openCount int, closeCount int, requiredLength int, chars []rune, cIndex int, seen map[string]bool) {
    if i >= len(s) && openCount == closeCount && cIndex == requiredLength {
        seen[string(chars)] = true
        return
    }
    
    if closeCount > openCount || cIndex > requiredLength || i >= len(s) {
        return
    }
    
    c := rune(s[i])
    if c == '(' {
        if cIndex < requiredLength {
            chars[cIndex] = c
            helper(s, i+1, openCount+1, closeCount, requiredLength, chars, cIndex+1, seen) 
        }
        helper(s, i+1, openCount, closeCount, requiredLength, chars, cIndex, seen)
    } else if c == ')' {
        if cIndex < requiredLength {
            chars[cIndex] = c
            helper(s, i+1, openCount, closeCount+1, requiredLength, chars, cIndex+1, seen)   
        }
        helper(s, i+1, openCount, closeCount, requiredLength, chars, cIndex, seen)  
    } else {
        if cIndex < requiredLength {
            chars[cIndex] = c
            helper(s, i+1, openCount, closeCount, requiredLength, chars, cIndex+1, seen)
        }    
    }
    
    

}

func requiredLength(s string) int {
    curCloseLeft := 0
    maxCloseLeft := 0
    curOpenRight := 0
    maxOpenRight := 0
    for i := 0; i < len(s); i++ {
        first := string(s[i])
        if first == ")" {
            curCloseLeft++
        } else if first == "(" {
            curCloseLeft--
        }
        if curCloseLeft > maxCloseLeft {
            maxCloseLeft = curCloseLeft
        }
        
        second := string(s[len(s)-1-i])
        if second == "(" {
            curOpenRight++
        } else if second == ")" {
            curOpenRight--
        }
        if curOpenRight > maxOpenRight {
            maxOpenRight = curOpenRight
        }
    }
    
    return len(s) - maxOpenRight - maxCloseLeft
}

