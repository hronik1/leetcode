import "strconv"

func evalRPN(tokens []string) int {
    stack := []int{}
    for _, a := range tokens {
        if i, err := strconv.Atoi(a); err == nil {
            stack = append(stack, i)
            continue
        }

        first := stack[len(stack)-2]
        second := stack[len(stack)-1]
        stack = stack[:len(stack)-2]

        val := 0
        if a == "+" {
            val = first + second
        } else if a == "-" {
            val = first - second
        } else if a == "/" {
            val = first / second
        } else if a == "*" {
            val = first * second
        } else {
            //err
        }

        stack = append(stack, val)
    }

    return stack[0]
}
