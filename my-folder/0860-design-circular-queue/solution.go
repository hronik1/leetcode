type MyCircularQueue struct {
    headI int
    tailI int
    length int
    capacity int
    contents []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        headI: -1,
        tailI: -1,
        length: 0,
        capacity: k,
        contents: make([]int, k),
    }
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.length >= this.capacity {
        return false
    }
    
    if this.length == 0 {
        this.headI, this.tailI = 0, 0
    } else {
        this.tailI = (this.tailI+1)%this.capacity
    }
    
    this.contents[this.tailI] = value
    this.length++
    
    return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
    if this.length == 0 {
        return false
    }
    
    if this.length == 1 {
        this.headI, this.tailI = -1, -1
    } else {
        this.headI = (this.headI+1)%this.capacity
    }
    
    this.length--
    return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
    if this.length == 0 {
        return -1
    }
    
    return this.contents[this.headI]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
    if this.length == 0 {
        return -1
    }
    
    return this.contents[this.tailI]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
    return this.length == 0
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
    return this.length == this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
