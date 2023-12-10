type RecentCounter struct {
    pings []int
}


func Constructor() RecentCounter {
    return RecentCounter{pings: []int{}}
}


func (this *RecentCounter) Ping(t int) int {
    this.pings = append(this.pings, t)
    newStart := 0
    for ; this.pings[newStart] < t - 3000; newStart++ {}
    this.pings = this.pings[newStart:]
    return len(this.pings)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
