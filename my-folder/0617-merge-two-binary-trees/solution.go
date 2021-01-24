/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    
    t := TreeNode{}
    var l1 *TreeNode
    var r1 *TreeNode
    if t1 != nil {
        t.Val += t1.Val
        l1 = t1.Left
        r1 = t1.Right
    }
    
    var l2 *TreeNode
    var r2 *TreeNode
    if t2 != nil {
        t.Val += t2.Val
        l2 = t2.Left
        r2 = t2.Right
    }
    
    t.Left = mergeTrees(l1, l2)
    t.Right = mergeTrees(r1, r2)
    
    return &t
}
