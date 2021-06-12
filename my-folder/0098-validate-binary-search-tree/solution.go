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
     
    rootOut := helper(root)
    return rootOut.IsValid
}

type HelperOut struct {
    IsValid bool
    Min *int
    Max *int
}

func helper(root *TreeNode) HelperOut {
    if root == nil {
        return HelperOut{IsValid: true}
    }
    leftOut := helper(root.Left)
    rightOut := helper(root.Right)
    
    out := HelperOut{IsValid: false}
    if !leftOut.IsValid || !rightOut.IsValid {
        return out
    }
    
    if leftOut.Max != nil && *leftOut.Max >= root.Val {
        return out
    }
    
    if rightOut.Min != nil && *rightOut.Min <= root.Val {
        return out
    }
    
    out = HelperOut{
        IsValid: true,
        Min: &root.Val,
        Max: &root.Val,
    }
    
    if leftOut.Min != nil {
        out.Min = leftOut.Min
    }
    if rightOut.Max != nil {
        out.Max = rightOut.Max
    }
    
    return out
}

