import "slices"
func groupAnagrams(strs []string) [][]string {
    mapping := map[string][]string{}
    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        s := string(runes)
        
        if l, ok := mapping[s]; ok {
            mapping[s] = append(l, str)
        } else {
            mapping[s] = []string{str}
        }
    }

    out := [][]string{}
    for _, v := range mapping {
        out = append(out, v)
    }

    return out
}
