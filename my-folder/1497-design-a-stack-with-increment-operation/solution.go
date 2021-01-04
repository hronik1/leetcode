type CustomStack struct {
    Capacity int
    Items []int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack {
        Capacity: maxSize,
        Items: []int{},
    }
}


func (this *CustomStack) Push(x int)  {
    if len(this.Items) >= this.Capacity {
        return
    }
    
    this.Items = append(this.Items, x)
}


func (this *CustomStack) Pop() int {
    if len(this.Items) == 0 {
        return -1
    }
    
    out := this.Items[len(this.Items)-1]
    this.Items = this.Items[:len(this.Items)-1]
    
    return out
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; i < len(this.Items) && i < k; i++ {
        this.Items[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
