import "strconv"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return strings.Join(buildStrings(root), ",")
}

func buildStrings(root *TreeNode) []string {
    if root == nil {
        return []string{""}
    }
    
    out := []string{strconv.Itoa(root.Val)}
    
    left := buildStrings(root.Left)
    out = append(out, left...)
    
    right := buildStrings(root.Right)
    out = append(out, right...)
    
    return out
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    split := strings.Split(data, ",")
    i := 0
    return buildTree(split, &i)
}

func buildTree(split []string, i *int) *TreeNode {
    if *i >= len(split) {
        return nil
    }
    
    
    if split[*i] == "" {
        *i++
        return nil
    }
    
    v, _ := strconv.Atoi(split[*i])
    out := TreeNode{
        Val: v,
    }
    
    *i++
    out.Left = buildTree(split, i)
    out.Right = buildTree(split, i)
    
    return &out
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
