var vowels = map[byte]bool {
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func maxVowels(s string, k int) int {
    count := 0
    for i := 0; i < k; i += 1 {
        if _, ok := vowels[s[i]]; ok {
            count += 1
        }
    }
    
    maxCount := count
    for prev, next := 0, k; next < len(s); prev, next = prev+1, next+1 {
        if _, ok := vowels[s[prev]]; ok {
            count -= 1
        }

        if _, ok := vowels[s[next]]; ok {
            count += 1
        }

        if count > maxCount {
            maxCount = count
        }
    }

    return maxCount
    
}
