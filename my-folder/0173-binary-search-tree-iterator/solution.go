/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    Nodes []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    nodes := []*TreeNode{root}
    cur := root.Left
    for cur != nil {
        nodes = append(nodes, cur)
        cur = cur.Left 
    }
    
    return BSTIterator{
        Nodes: nodes,
    }
}


func (this *BSTIterator) Next() int {
    l := len(this.Nodes)
    node := this.Nodes[l-1]
    this.Nodes = this.Nodes[:l-1]
    if node.Right != nil {
        this.Nodes = append(this.Nodes, node.Right)
        cur := node.Right.Left
        for cur != nil {
            this.Nodes = append(this.Nodes, cur)
            cur = cur.Left
        }
    }
    
    return node.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.Nodes) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
