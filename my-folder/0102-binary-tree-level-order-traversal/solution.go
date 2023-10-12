/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    
    curLevel := []*TreeNode{root}
    curValues := []int{}
    nextLevel := []*TreeNode{}
    for len(curLevel) > 0 {
        cur := curLevel[0]
        curValues = append(curValues, cur.Val)
        
        if cur.Left != nil {
            nextLevel = append(nextLevel, cur.Left)
        }
        
        if cur.Right != nil {
            nextLevel = append(nextLevel, cur.Right)
        }
        
        if len(curLevel) == 1 {
            curLevel = nextLevel
            res = append(res, curValues)
            curValues = []int{}
            nextLevel = []*TreeNode{}
        } else {
            curLevel = curLevel[1:]
        }
    }
    
    return res
}
