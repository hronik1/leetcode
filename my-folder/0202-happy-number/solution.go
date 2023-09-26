func isHappy(n int) bool {
    seen := map[int]bool{}
    num := n
    for num != 1 {
        if _, ok := seen[num]; ok {
            return false
        } else {
            seen[num] = true
        }

        num = squareDigits(num)
    }

    return true
}

func squareDigits(num int) int {
    sum := 0
    for num > 0 {
        lastDigit := num%10
        sum += (lastDigit*lastDigit)
        num /= 10
    }

    return sum
}
