/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    res = helper(root, res)
    return res
}

func helper(root *TreeNode, res []int) []int {
    if root == nil {
        return res
    }
    
    res = helper(root.Left, res)
    res = append(res, root.Val)
    res = helper(root.Right, res)
    
    return res
}
