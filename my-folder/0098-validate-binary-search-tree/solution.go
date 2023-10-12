/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    if root.Left != nil {
        maxLeftVal := maxVal(root.Left)
        if maxLeftVal >= root.Val || !isValidBST(root.Left) {
            return false
        }
    }
    
    if root.Right != nil {
        minRightVal := minVal(root.Right)
        if minRightVal <= root.Val || !isValidBST(root.Right) {
            return false
        }
    }
    
    return true
}

func maxVal(root *TreeNode) int {
    res := root.Val
    if root.Left != nil {
        leftVal := maxVal(root.Left)
        if leftVal > res {
            res = leftVal
        }
    }
    
    if root.Right != nil {
        rightVal := maxVal(root.Right)
        if rightVal > res {
            res = rightVal
        }
    }
    
    return res
}

func minVal(root *TreeNode) int {
    res := root.Val
    if root.Left != nil {
        leftVal := minVal(root.Left)
        if leftVal < res {
            res = leftVal
        }
    }
    
    if root.Right != nil {
        rightVal := minVal(root.Right)
        if rightVal < res {
            res = rightVal
        }
    }
    
    return res
}
