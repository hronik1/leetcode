import "fmt"
import "strconv"
import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    if root.Left == nil && root.Right == nil {
        return root.Val
    }

    sum := 0
    allDigits := generateAllDigits(root)
    for _, digits := range allDigits {
        slices.Reverse(digits)
        sum += toInt(digits)
    }

    return sum
}

func toInt(digits []int) int {
    var b strings.Builder
    for _, digit := range digits {
        fmt.Fprintf(&b, "%d", digit)
    }

    s := b.String()
    val, _ := strconv.Atoi(s)
    
    return val
}

func generateAllDigits(root *TreeNode) [][]int {
    out := [][]int{}
    if root == nil {
        return out
    }

    if root.Left == nil && root.Right == nil {
        out = append(out, []int{root.Val})
        return out
    }

    leftDigits := generateAllDigits(root.Left)
    rightDigits := generateAllDigits(root.Right)
    for _, digits := range leftDigits {
        digits = append(digits, root.Val)
        out = append(out, digits)
    }

    for _, digits := range rightDigits {
        digits = append(digits, root.Val)
        out = append(out, digits)
    }

    return out
}
