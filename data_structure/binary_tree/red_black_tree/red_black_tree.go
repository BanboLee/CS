package red_black_tree

const (
	RED byte = iota
	BLACK
)

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
	color  byte
}

type RedBlackTree struct {
	root *Node
	n    int
}

func (rbt *RedBlackTree) rotateLeft(u *Node) *Node {
	// 左旋将右孩子提上去
	child := u.Right
	child.Parent = u.Parent
	if child.Parent != nil {
		if child.Parent.Left == u {
			child.Parent.Left = child
		} else {
			child.Parent.Right = child
		}
	}

	u.Right = child.Left
	if u.Right != nil {
		u.Right.Parent = u
	}
	u.Parent = child
	child.Left = u

	// 变成根节点
	if u == rbt.root {
		rbt.root = child
		rbt.root.Parent = nil
	}
	return child
}

func (rbt *RedBlackTree) rotateRight(u *Node) *Node {
	// 右旋，将左孩子提上去
	leftChild := u.Left
	leftChild.Parent = u.Parent
	if leftChild.Parent != nil {
		if leftChild.Parent.Left == u {
			leftChild.Parent.Left = leftChild
		} else {
			leftChild.Parent.Right = leftChild
		}
	}

	u.Left = leftChild.Right
	if u.Left != nil {
		u.Left.Parent = u
	}
	u.Parent = leftChild
	leftChild.Right = u
	if u == rbt.root {
		rbt.root = leftChild
		rbt.root.Parent = nil
	}
	return leftChild
}

// colour u red and its two children black.
func pushBlack(u *Node) {
	u.color--
	u.Left.color++
	u.Right.color++
}

// reverse pushBlack.
func pullBlack(u *Node) {
	u.color++
	u.Left.color--
	u.Right.color--
}

func (rbt *RedBlackTree) findPre(node *Node) *Node {
	var pre *Node = nil
	cur := rbt.root
	for cur != nil {
		pre = cur
		if node.Val < cur.Val {
			cur = cur.Left
		} else if node.Val > cur.Val {
			cur = cur.Right
		} else {
			return node
		}
	}
	return pre
}

func (rbt *RedBlackTree) bstAdd(pre, node *Node) bool {
	if pre == nil {
		rbt.root = node
		return true
	} else {
		if node.Val < pre.Val {
			pre.Left = node
		} else if node.Val > pre.Val {
			pre.Right = node
		} else {
			return false
		}
		node.Parent = pre
	}
	return true
}
