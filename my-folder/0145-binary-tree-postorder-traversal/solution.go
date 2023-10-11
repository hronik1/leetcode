/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    postorderHelper(root, &res)
    
    return res
}

func postorderHelper(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    
    postorderHelper(root.Left, res)
    postorderHelper(root.Right, res)
    *res = append(*res, root.Val)
}
