func kthGrammar(n int, k int) int {
    if n == 1 {
        return 0
    }
    
    rowLength := int(math.Pow(2.0, float64(n-1)))
    if k > rowLength/2 {
        prev := kthGrammar(n-1, k-rowLength/2)
        if prev == 1 {
            return 0
        } else {
            return 1
        }
    } else {
        return kthGrammar(n-1, k)
    }
}

// func genNthRow(n int) []int {
//     if n == 1 {
//         return []int{0}
//     }
    
//     prev := genNthRow(n-1)
//     out := []int{}
//     for _, v := range prev {
//         if v == 0 {
//             out = append(out, 0)
//             out = append(out, 1)
//         } else {
//             out = append(out, 1)
//             out = append(out, 0)
//         }
//     }
    
//     return out
// }
