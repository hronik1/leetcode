func snakesAndLadders(board [][]int) int {
    visited := map[int] bool{1:true}
    candidates := []int{1}
    moves := 0
    size := len(board)*len(board)
    length := len(board)
    for {
        newCandidates := []int{}
        moves++
        for _, candidate := range candidates {
            for offset := 1; offset <= 6; offset++ {
                newCandidate := candidate + offset
                
                // check for shoot or ladder
                resp := to2D(newCandidate, length)
                if resp.I < 0 || resp.J < 0 {
                    continue
                }
                
                if destination := board[resp.I][resp.J]; destination > 0 {
                    if destination >= size {
                        return moves
                    }
                    
                    if _, destinationVisited := visited[destination]; !destinationVisited {
                        visited[destination] = true
                        newCandidates = append(newCandidates, destination)
                    }
                } else {
                    if newCandidate >= size {
                        return moves
                    }
                    
                    if _, newCandidateVisited := visited[newCandidate]; !newCandidateVisited {
                        visited[newCandidate] = true
                        newCandidates = append(newCandidates, newCandidate)  
                    }   
                }
            }
        }
        
        if len(newCandidates) == 0 {
            break
        }
        
        candidates = newCandidates
    }
    
    return -1
}

type to2DResponse struct {
    I int
    J int
}

func to2D(square int, length int) to2DResponse {
    i := (length-1) - (square-1)/length
    //even rows from bottom start left, odd rows from bottom start from right
    j := (square-1)%length
    if (length-i)%2 == 0 {
        j = length-1-j
    }
    
    return to2DResponse {
        I: i,
        J: j,
    }
} 
