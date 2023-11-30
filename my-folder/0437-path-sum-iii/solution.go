/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    count := 0
    level := []*TreeNode{root}
    for len(level) > 0 {
        nextLevel := []*TreeNode{}
        for _, node := range level {
            count += pathSumHelper(node, targetSum, 0)
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }

        level = nextLevel
    }

    return count
}

func pathSumHelper(root *TreeNode, targetSum int, currentSum int) int {
    if root == nil {
        return 0
    }

    //fmt.Printf("val: %d currentSum: %d\n", root.Val, currentSum)
    count := 0
    currentSum += root.Val
    if currentSum == targetSum {
        fmt.Println("found")
        count += 1
    }

    count += pathSumHelper(root.Left, targetSum, currentSum)
    count += pathSumHelper(root.Right, targetSum, currentSum)

    return count
}
