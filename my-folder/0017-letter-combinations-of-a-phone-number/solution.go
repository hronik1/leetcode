var mappings = map[rune][]rune{
    '2': []rune{'a', 'b', 'c'},
    '3': []rune{'d', 'e', 'f'},
    '4': []rune{'g', 'h', 'i'},
    '5': []rune{'j', 'k', 'l'},
    '6': []rune{'m', 'n', 'o'},
    '7': []rune{'p', 'q', 'r', 's'},
    '8': []rune{'t', 'u', 'v'},
    '9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    out := []string {}
    if len(digits) == 0 {
        return out
    }

    if len(digits) == 1 {
        for _, v := range mappings[rune(digits[0])] {
            out = append(out, string(v))
        }

        return out
    }

    nextCombos := letterCombinations(digits[1:])
    for _, v := range mappings[rune(digits[0])] {
        for _, combo := range nextCombos {
            out = append(out, string(v) + combo)
        }
    }

    return out
}
