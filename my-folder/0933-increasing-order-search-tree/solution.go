/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    out := increasingBSTHelper(root)
    return out.root
}

type increasingBSTHelperResult struct {
    root *TreeNode
    rightMost *TreeNode
}

func increasingBSTHelper(root *TreeNode) increasingBSTHelperResult {
    out := increasingBSTHelperResult{
        root: root,
        rightMost: root,
    }
    
    if root.Right != nil {
        right := increasingBSTHelper(root.Right)
        out.rightMost = right.rightMost
        root.Right = right.root
    }
    
    if root.Left != nil {
        left := increasingBSTHelper(root.Left)
        left.rightMost.Right = root
        out.root = left.root
    }
    
    root.Left = nil
    
    return out
}
