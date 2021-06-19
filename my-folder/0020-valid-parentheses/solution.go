func isValid(s string) bool {
    stack := []rune{}
    for _, v := range s {
        v = rune(v)
        if v == '(' || v == '{' || v == '[' {
            stack = append(stack, v)
            continue
        }
        
        if len(stack) == 0 {
            return false
        }
        
        if v == ')' {
            if stack[len(stack)-1] != '(' {
                return false
            }
        } else if v == ']' {
            if stack[len(stack)-1] != '[' {
                return false
            }
        } else if v == '}' {
            if stack[len(stack)-1] != '{' {
                return false
            }
        }
        
        stack = stack[:len(stack)-1]
    }
    
    return len(stack) == 0
}
