import "strings"

func predictPartyVictory(senate string) string {
    votes := []rune{}
    prevSenate := senate
    for {
        var newSenateBuilder strings.Builder
        for _, v := range prevSenate {
            if len(votes) == 0 || v == votes[0] {
                votes = append(votes, v)
                newSenateBuilder.WriteString(string(v))
            }  else {
                votes = votes[:len(votes) - 1]
            }
        }

        newSenate := newSenateBuilder.String()
        if len(prevSenate) == len(newSenate) {
            break
        }
        
        prevSenate = newSenate
    }

    if votes[0] == 'R' {
        return "Radiant"
    }
    return "Dire"

}
