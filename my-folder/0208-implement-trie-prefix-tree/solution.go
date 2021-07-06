type Node struct {
    Children map[rune]*Node
    Val string
}

type Trie struct {
    Root *Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        Root: &Node{Children: map[rune]*Node{}},
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
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


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.Root
    for _, r := range word {
        if child, ok := node.Children[r]; ok {                
            node = child
        } else {
            return false
        }    
    }
    
    return node.Val != ""
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this.Root
    for _, r := range prefix {
        if child, ok := node.Children[r]; ok {            
            node = child   
        } else {
            return false
        }    
    }
    
    return true    
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
