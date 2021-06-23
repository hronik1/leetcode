/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    helper(root)
}

func helper(root *TreeNode) HelperOut {
    if root == nil {
        return HelperOut{}
    }
    
    var left HelperOut
    if root.Left != nil {
        left = helper(root.Left)
    }
    
    var right HelperOut 
    if root.Right != nil {
        right = helper(root.Right)
    }
    
    out := HelperOut{
        Head: root,
        Tail: root,
    }
    if left.Head != nil {
        out.Tail.Right = left.Head
        out.Tail = left.Tail
    }
    
    if right.Head != nil {
        out.Tail.Right = right.Head
        out.Tail = right.Tail
    }
    
    root.Left = nil
    return out
}

type HelperOut struct {
    Head *TreeNode
    Tail *TreeNode
}
