/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths (root *TreeNode) int {
    return helper(root, map[int]int{}, map[int]int{})
}

func helper(root *TreeNode, evenCounts map[int]int, oddCounts map[int]int) int {
    if root == nil {
        return 0
    }
    
    valCount := 1
    if evenCount, ok := evenCounts[root.Val]; ok {
        valCount += evenCount
        delete(evenCounts, root.Val)
    }
    
    if oddCount, ok := oddCounts[root.Val]; ok {
        valCount += oddCount
        delete(oddCounts, root.Val)
    }
    
    if valCount%2 == 0 {
        evenCounts[root.Val] = valCount
    } else {
        oddCounts[root.Val] = valCount
    }
     
    pathCounts := 0
    if root.Left == nil && root.Right == nil {
        if len(oddCounts) == 0 || len(oddCounts) == 1 {
            pathCounts++
        }
    } else {
        pathCounts += helper(root.Left, evenCounts, oddCounts)
        pathCounts += helper(root.Right, evenCounts, oddCounts)
    }
    
    valCount = -1
    if evenCount, ok := evenCounts[root.Val]; ok {
        valCount += evenCount
        delete(evenCounts, root.Val)
    }
    
    if oddCount, ok := oddCounts[root.Val]; ok {
        valCount += oddCount
        delete(oddCounts, root.Val)
    }
    
    if valCount%2 == 1 {
        oddCounts[root.Val] = valCount
    } else if valCount > 0 {
        evenCounts[root.Val] = valCount
    }
    
    return pathCounts
}
