/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return isSymmetricHelper([]*TreeNode{root})
}

func isSymmetricHelper(level []*TreeNode) bool {
    onlyNils := true
    nextLevel := []*TreeNode{}
    for i, node := range level {
        mirror := level[len(level)-1-i]
        if node != nil {
            onlyNils = false
            if mirror == nil || node.Val != mirror.Val {
                return false
            }
            
            nextLevel = append(nextLevel, node.Left)
            nextLevel = append(nextLevel, node.Right)
        }
    }
    
    if onlyNils {
        return true
    }
    
    return isSymmetricHelper(nextLevel)
}
