type List struct {
    Head *Node
    Tail *Node
    Size int
}

func (l *List) InsertHead(n *Node) {
    if n == nil {
        return
    }
    
    if l.Size == 0 {
        l.Head = n
        l.Tail = n
    } else {
        n.Next = l.Head
        n.Next.Prev = n
        l.Head = n
    }
    
    l.Size++
}

func (l *List) Remove(n *Node) {
    if n == nil {
        return
    }
    
    if l.Tail == n {
        l.Tail = n.Prev
    }
    if l.Head == n {
        l.Head = n.Next
    }
    
    if n.Prev != nil {
        n.Prev.Next = n.Next
    }
    if n.Next != nil {
        n.Next.Prev = n.Prev
    }
    
    n.Prev, n.Next = nil, nil
    l.Size--
}

type Node struct {
    Key int
    Val int
    Prev *Node
    Next *Node
}

type LRUCache struct {
    Capacity int
    Index map[int]*Node
    Items *List
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        Capacity: capacity,
        Index: map[int]*Node{},
        Items: &List{},
    }
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.Index[key]; ok {
        this.Items.Remove(node)
        this.Items.InsertHead(node)
        return node.Val
    }
    
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.Index[key]; ok {
        this.Items.Remove(node)
        node.Val = value
        this.Items.InsertHead(node)
        return
    }
    
    if this.Items.Size >= this.Capacity {
        delete(this.Index, this.Items.Tail.Key)
        this.Items.Remove(this.Items.Tail)
    }
    
    node := &Node{Key: key, Val: value}
    this.Items.InsertHead(node)
    this.Index[key] = node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
