func asteroidCollision(asteroids []int) []int {
    out := []int{}
    for _, v := range asteroids {
        if len(out) == 0 || v > 0 || out[len(out)-1] < 0 {
            out = append(out, v)
        } else {
            abs := -1 * v
            for i := len(out) - 1; i >= 0 && out[i] > 0 && abs > out[i]; i-- {
                out = out[:i]
            }

            if len(out) == 0 || out[len(out) - 1] < 0{
                out = append(out, v)
            } else if abs == out[len(out) - 1] {
                out = out[:len(out)-1]
            }
        }
    }

    return out
}
