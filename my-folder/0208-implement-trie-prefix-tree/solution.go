type Trie struct {
    children map[rune]*Trie
}


func Constructor() Trie {
    return Trie{children: make(map[rune]*Trie)}
}


func (this *Trie) Insert(word string)  {
    cur := this
    for _, r := range word {
        next, ok := cur.children[r]
        if !ok {
            n := Constructor()
            next = &n
            cur.children[r] = next
        }

        cur = next
    }

    cur.children['*'] = nil
}


func (this *Trie) Search(word string) bool {
    cur := this
    for _, r := range word {
        next, ok := cur.children[r]
        if !ok {
            return false
        }

        cur = next
    }

    _, ok := cur.children['*']
    return ok
}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this
    for _, r := range prefix {
        next, ok := cur.children[r]
        if !ok {
            return false
        }

        cur = next
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
