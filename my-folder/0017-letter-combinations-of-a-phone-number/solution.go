var letters = map[string][]string {
    "2": []string{"a","b","c"},
    "3": []string{"d","e","f"},
    "4": []string{"g","h","i"},
    "5": []string{"j","k","l"},
    "6": []string{"m","n","o"},
    "7": []string{"p","q","r","s"},
    "8": []string{"t","u","v"},
    "9": []string{"w","x","y","z"},
}


func letterCombinations(digits string) []string {
    out := []string{}
    if len(digits) == 0 {
        return out
    }
    
    if len(digits) == 1 {
        return letters[digits]
    }
    
    subStrings := letterCombinations(digits[1:])
    for _, letter := range letters[string(digits[0])] {
        for _, subString := range subStrings {
            out = append(out, letter+subString)
        }
    }
    
    return out
}
