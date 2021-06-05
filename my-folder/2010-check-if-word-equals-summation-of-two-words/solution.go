func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return wordToInt(firstWord) + wordToInt(secondWord) == wordToInt(targetWord)
}

func wordToInt(word string) int {
    sum := 0
    l := len(word)
    pow := 1
    for i := 0; i < l; i++ {
        sum += pow*int((word[l-1-i] - 'a'))
        pow *= 10
    }
    
    return sum
}
