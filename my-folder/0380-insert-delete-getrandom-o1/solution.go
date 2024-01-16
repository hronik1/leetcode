type RandomizedSet struct {
    index map[int]int
    orderedItems []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        index: map[int]int{},
        orderedItems: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.index[val]; ok {
        return false
    }

    this.orderedItems = append(this.orderedItems, val)
    this.index[val] = len(this.orderedItems)-1

    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    i, ok := this.index[val]
    if !ok {
        return false
    }

    lastItem := this.orderedItems[len(this.orderedItems) - 1]
    this.orderedItems[i] = lastItem
    this.index[lastItem] = i
    
    delete(this.index, val)
    this.orderedItems = this.orderedItems[:len(this.orderedItems) - 1] 

    return true
}


func (this *RandomizedSet) GetRandom() int {
    if len(this.orderedItems) == 0 {
        return -1
    }

    randI := rand.Int31n(int32(len(this.orderedItems)))
    return this.orderedItems[randI] 
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
