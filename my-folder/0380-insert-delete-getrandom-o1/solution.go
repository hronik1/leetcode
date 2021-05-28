import "math/rand"

type RandomizedSet struct {
    s map[int]int
    l []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet {
        s: map[int]int{},
        l: []int{},
    }
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.s[val]; ok {
        return false
    }
    
    this.s[val] = len(this.l)
    this.l = append(this.l, val)
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if index, ok := this.s[val]; ok {
        if index != len(this.l) - 1 {
            temp := this.l[len(this.l) - 1]
            this.s[temp] = index
            this.l[index] = temp
        }
        delete(this.s, val)
        this.l = this.l[:len(this.l) - 1]
        return true
    }
    
    return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    randI := rand.Intn(len(this.l))
    return this.l[randI]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
