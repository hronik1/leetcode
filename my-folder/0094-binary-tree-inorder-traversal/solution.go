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
    inorderHelper(root, &res)
    
    return res
}

func inorderHelper(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    
    inorderHelper(root.Left, res)
    *res = append(*res, root.Val)
    inorderHelper(root.Right, res)
}
