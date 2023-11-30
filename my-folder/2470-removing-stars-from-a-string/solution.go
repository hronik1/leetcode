func removeStars(s string) string {
    out := []rune{}
    for _, v := range s {
        if v == '*' {
            if len(out) > 0 {
                out = out[:len(out)-1]
            }
        } else {
            out = append(out, v)
        }
    }

    return string(out)
}
