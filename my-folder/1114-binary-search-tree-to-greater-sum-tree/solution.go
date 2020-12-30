/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    _ = gstHelper(root, 0)
    
    return root
}

func gstHelper(root *TreeNode, greaterThan int) int {
    if root == nil {
        return 0
    }
    
    sum := root.Val
    sum += gstHelper(root.Right, greaterThan)
    root.Val = sum + greaterThan
    sum += gstHelper(root.Left, root.Val)
    
    return sum
}
