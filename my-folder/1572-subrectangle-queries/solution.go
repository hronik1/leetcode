type SubrectangleQueries struct {
    Vals [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{
        Vals: rectangle,
    }
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for row := row1; row <= row2; row++ {
        for col := col1; col <= col2; col++ {
            this.Vals[row][col] = newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.Vals[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
