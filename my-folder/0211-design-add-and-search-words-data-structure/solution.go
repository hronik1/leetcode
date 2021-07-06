type Node struct {
    Children map[rune]*Node
    Val string
}

type WordDictionary struct {
    Root *Node
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{
        Root: &Node{
            Children: map[rune]*Node{},
        },
    }
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.Root
    for _, r := range word {
        var ok bool
        var child *Node
        if child, ok = node.Children[r]; !ok {
            child = &Node{
                Children: map[rune]*Node{},
            }
            node.Children[r] = child
        }
        
        node = child
    }
    
    node.Val = word
}


func (this *WordDictionary) Search(word string) bool {
    candidates := []*Node{this.Root}
    for _, r := range word {
        if len(candidates) == 0 {
            return false
        }
        
        newCandidates := []*Node{}
        for _, node := range candidates {
            if r == '.' {
                for _, child := range node.Children {
                    newCandidates = append(newCandidates, child)
                }
            } else {
                if child, ok := node.Children[r]; ok {
                    newCandidates = append(newCandidates, child)
                }
            }
        }
        
        candidates = newCandidates
    }
    
    for _, node := range candidates {
        if node.Val != "" {
            return true
        }
    }
    
    return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
