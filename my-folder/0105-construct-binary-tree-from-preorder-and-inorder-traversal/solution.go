var preIndex int 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderIndex := map[int]int{}
    for k, v := range inorder {
        inorderIndex[v] = k
    }
    
    preIndex = 0
    return helper(preorder, inorder, inorderIndex, 0, len(inorder)-1)
}

func helper(preorder []int, inorder []int, inorderIndex map[int]int, lo int, hi int) *TreeNode {
    fmt.Println("lo: %d hi: %d", lo, hi)
    if lo > hi {
        return nil
    }
    
    rootVal := preorder[preIndex]
    preIndex++
    out := TreeNode{
        Val: rootVal,
    }
    
    l := inorderIndex[rootVal]-lo
    fmt.Println("left lo: %d hi: %d", inorderIndex[rootVal], inorderIndex[rootVal]+l-1)
    out.Left = helper(preorder, inorder, inorderIndex, lo, inorderIndex[rootVal]-1)
    fmt.Println("Right lo: %d hi: %d", inorderIndex[rootVal]+1, hi)
    out.Right = helper(preorder, inorder, inorderIndex, inorderIndex[rootVal]+1, hi)
    
    return &out
}
