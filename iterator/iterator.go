package iterator

type BinaryTree struct {
    Index     int
    LeftNode  *BinaryTree
    RightNode *BinaryTree
}

func InitTree(index int) *BinaryTree {
    return &BinaryTree{Index: index}
}

func (t *BinaryTree) LeftAdd(index int) {
    node := &BinaryTree{Index: index}
    t.LeftNode = node
}

func (t *BinaryTree) RightAdd(index int) {
    node := &BinaryTree{Index: index}
    t.RightNode = node
}

func (t *BinaryTree) BreadthFirstSearch() []int {
    breadthArr := []*BinaryTree{t}
    result := []int{}
    for len(breadthArr) > 0 {
        node := breadthArr[0]
        result = append(result, node.Index)
        if node.LeftNode != nil {
            breadthArr = append(breadthArr, node.LeftNode)
        }
        if node.RightNode != nil {
            breadthArr = append(breadthArr, node.RightNode)
        }
        breadthArr = breadthArr[1:]
    }
    return result
}

func (t *BinaryTree) DepthFirstSearch() []int {
    depthStack := []*BinaryTree{t}
    result := []int{}
    for len(depthStack) > 0 {
        node := depthStack[0]
        depthStack = depthStack[1:]
        result = append(result, node.Index)
        if node.RightNode != nil {
            depthStack = append([]*BinaryTree{node.RightNode}, depthStack...)
        }
        if node.LeftNode != nil {
            depthStack = append([]*BinaryTree{node.LeftNode}, depthStack...)
        }
    }
    return result
}
