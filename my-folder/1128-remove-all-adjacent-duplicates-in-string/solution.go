type RuneNode struct {
    Val rune
    Prev *RuneNode
    Next *RuneNode
}

func (r *RuneNode) Remove() {
    if r.Prev != nil {
        r.Prev.Next = r.Next
    }
    
    if r.Next != nil {
        r.Next.Prev = r.Prev
    }
}

func fromString(s string) *RuneNode {
    if len(s) == 0 {
        return nil
    }
    
    head := &RuneNode{Val:rune(s[0])}
    prev := head
    for i := 1; i < len(s); i++ {
        cur := &RuneNode{Val:rune(s[i])}
        cur.Prev = prev
        prev.Next = cur
        prev = cur
    }
    
    return head
}

func toString(h *RuneNode) string {
    var builder strings.Builder
    cur := h
    for cur != nil {
        builder.WriteRune(cur.Val)
        cur = cur.Next
    }
    
    return builder.String()
}

func removeDuplicates(s string) string {
    if len(s) == 0 {
        return s
    }
    
    cur := fromString(s)
    head := cur
    for cur != nil && cur.Next != nil { 
        if cur.Val == cur.Next.Val {
            cur.Remove()
            cur.Next.Remove()
            if cur == head {
                cur = cur.Next.Next
                head = cur
            } else {
                cur = cur.Prev
            }
        } else {
            cur = cur.Next
        }
    }
    
    return toString(head)
}
