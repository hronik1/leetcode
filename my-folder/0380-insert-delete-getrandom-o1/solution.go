type RandomizedSet struct {
    Index map[int]int
    Items []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
        Index: map[int]int{},
        Items: []int{},
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.Index[val]; ok {
        return false
    }
    
    l := len(this.Items)
    this.Index[val] = l
    this.Items = append(this.Items, val)
    
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if i, ok := this.Index[val]; !ok {
        return false
    } else {
        // swap this item with last item, as removing from the end of an end of a slice is easy
        if i != len(this.Items)-1 {
            lastVal := this.Items[len(this.Items)-1]
            this.Index[lastVal] = i
            this.Items[i] = lastVal
        }
        
        delete(this.Index, val)
        this.Items = this.Items[:len(this.Items)-1]
        
        return true
    }
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    return this.Items[rand.Intn(len(this.Items))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
