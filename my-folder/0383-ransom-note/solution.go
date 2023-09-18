func canConstruct(ransomNote string, magazine string) bool {
    noteCounts := runeCounts(ransomNote)
    magazineCounts := runeCounts(magazine)
    for r, noteCount := range noteCounts {
        if magazineCount, ok := magazineCounts[r]; !ok || magazineCount < noteCount{
            return false
        }
    }

    return true
}

func runeCounts(words string) map[rune]int {
    counts := make(map[rune]int)
    for _, r := range words {
        if count, ok := counts[r]; ok {
            counts[r] = count+1
        } else {
            counts[r] = 1
        }
    }

    return counts
}
