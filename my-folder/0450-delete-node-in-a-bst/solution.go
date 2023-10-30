/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if root.Val == key {
        return del(root)
    }
    
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        root.Right = deleteNode(root.Right, key)
    }
    
    return root
}

func del(root *TreeNode) *TreeNode {
    if root.Left == nil && root.Right == nil {
        return nil
    }
    
    if root.Left == nil { 
        return root.Right
    } else if root.Right == nil {
        return root.Left
    }
    
    // swap with leftmost descendant on right side
    var prevLeftMost *TreeNode
    leftMost := root.Right
    for leftMost.Left != nil {
        prevLeftMost = leftMost
        leftMost = leftMost.Left
    }
    
    root.Val = leftMost.Val
    if prevLeftMost == nil {
        root.Right = leftMost.Right
    } else {
        prevLeftMost.Left = leftMost.Right
    }
    
    return root
}
