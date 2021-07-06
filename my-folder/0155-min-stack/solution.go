type MinStackItem struct {
    Val int
    MinVal int
}

type MinStack struct {
    Items []MinStackItem
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        Items: []MinStackItem{},
    }
}


func (this *MinStack) Push(val int)  {
    item := MinStackItem{
        Val: val,
        MinVal: val,
    }
    
    l := len(this.Items)
    if l > 0 && this.Items[l-1].MinVal < val {
        item.MinVal = this.Items[l-1].MinVal
    }
    
    this.Items = append(this.Items, item)
}


func (this *MinStack) Pop()  {
    this.Items = this.Items[:len(this.Items)-1]
}


func (this *MinStack) Top() int {
    return this.Items[len(this.Items)-1].Val
}


func (this *MinStack) GetMin() int {
    return this.Items[len(this.Items)-1].MinVal
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
