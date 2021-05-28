func judgeCircle(moves string) bool {
    if len(moves)%2 != 0 {
        return false
    }
    
    x, y := 0, 0
    for _, c := range moves {
        if c == 'U' {
            y++
        } else if c == 'D' {
            y--
        } else if c == 'R' {
            x++
        } else if c == 'L' {
            x--
        }
    }
    
    return x == 0 && y == 0
}
