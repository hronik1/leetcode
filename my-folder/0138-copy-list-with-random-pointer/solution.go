/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToNewMapping := map[*Node]*Node{}

    var newHead *Node
    var newPrev *Node
    for cur := head; cur != nil; cur = cur.Next {
        newCur := &Node{Val: cur.Val}
        if newHead == nil {
            newHead = newCur
        } else {
            newPrev.Next = newCur
        }

        oldToNewMapping[cur] = newCur
        newPrev = newCur
    }

    for cur, newCur := head, newHead; cur != nil && newCur != nil; cur, newCur = cur.Next, newCur.Next {
        if cur.Random != nil {
            newCur.Random = oldToNewMapping[cur.Random]
        }
    } 

    return newHead
}
