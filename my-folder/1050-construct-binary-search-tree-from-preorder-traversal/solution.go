/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    return bst(preorder, 0, len(preorder)-1)
}

func bst(preorder []int, lo int, hi int) *TreeNode{
    if lo > hi {
        return nil
    }
    
    var i int
    for i = lo+1; i <= hi; i++ {
        if preorder[i] > preorder[lo] {
            break
        }
    } 
    
    return &TreeNode {
        Val: preorder[lo],
        Left: bst(preorder, lo+1, i-1),
        Right: bst(preorder, i, hi),
    }
}
