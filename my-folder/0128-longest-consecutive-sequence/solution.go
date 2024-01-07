func longestConsecutive(nums []int) int {
    numSet := map[int]bool{}
    for _, num := range nums {
        numSet[num] = true
    }

    bestSequenceLength := 0
    for _, num := range nums {
        if numSet[num-1] {
            continue
        }

        curSequenceLength := sequenceLength(num, numSet)
        if curSequenceLength > bestSequenceLength {
            bestSequenceLength = curSequenceLength
        }
    }

    return bestSequenceLength
}

func sequenceLength(num int, numSet map[int]bool) int {
    length := 0
    for i := num; numSet[i]; i++ {
        length++
    }
    return length
}
