type StockSpanner struct {
    priceStack []int
    spanStack []int
}


func Constructor() StockSpanner {
    return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
    span := 1
    for len(this.priceStack) > 0 && this.priceStack[len(this.priceStack)-1] <= price {
        this.priceStack = this.priceStack[:len(this.priceStack)-1]
        span += this.spanStack[len(this.spanStack)-1]
        this.spanStack = this.spanStack[:len(this.spanStack)-1]
    }

    this.spanStack = append(this.spanStack, span)
    this.priceStack = append(this.priceStack, price)

    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
