func groupAnagrams(strs []string) [][]string {
    groups := map[string][]string{}
    for _, str := range strs {
        s := []rune(str)
        sort.Sort(sortRunes(s))
        sorted := string(s)
        if l, ok := groups[sorted]; ok {
            groups[sorted] = append(l, str)
        } else {
            groups[sorted] = []string{str}
        }
    }
    
    out := [][]string{}
    for _, v := range groups {
        out = append(out, v)
    }
    
    return out
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Len() int{
   return len(s)
}
func (s sortRunes) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}
