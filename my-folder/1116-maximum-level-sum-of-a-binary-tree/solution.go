/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    curLevelNodes := []*TreeNode{root}
    curLevel := 1
    maxLevel := 1
    maxLevelSum := root.Val
    for len(curLevelNodes) > 0 {
        nextLevelNodes := []*TreeNode{}
        curLevelSum := 0
        for _, v := range curLevelNodes {
            curLevelSum += v.Val
            if v.Left != nil {
                nextLevelNodes = append(nextLevelNodes, v.Left)
            }
            if v.Right != nil {
                nextLevelNodes = append(nextLevelNodes, v.Right)
            }
        }

        if curLevelSum > maxLevelSum {
            maxLevel = curLevel
            maxLevelSum = curLevelSum 
        }

        curLevel++
        curLevelNodes = nextLevelNodes
    }

    return maxLevel
}
