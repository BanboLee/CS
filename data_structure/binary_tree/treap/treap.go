package binary_tree

import (
	"math/rand"
	"time"
)

type TreapNode struct {
	Val   int
	p     int // priority
	Left  *TreapNode
	Right *TreapNode
}

func newNode(x int) *TreapNode {
	rand.Seed(time.Now().UnixNano())
	return &TreapNode{
		Val: x,
		p:   rand.Int(),
	}
}

func Search(root *TreapNode, x int) *TreapNode {
	if root == nil || root.Val == x {
		return root
	}
	if x < root.Val {
		return Search(root.Left, x)
	} else {
		return Search(root.Right, x)
	}
}

/*
妙啊我草。
T1, T2 and T3 are subtrees of the tree rooted with y
(on left side) or x (on right side)
              y                               x
             / \     Right Rotation          /  \
            x   T3   – – – – – – – >        T1   y
           / \       < - - - - - - -            / \
          T1  T2     Left Rotation            T2  T3
*/
func leftRotation(root *TreapNode) *TreapNode {
	// 左旋， 让右孩子上去
	child := root.Right
	// child的左子树要变更到父节点右子树下
	w := child.Left
	// 右孩子上去以后左子树为父节点
	child.Left = root
	// 左子树右节点原来为child， 现在变成了child的左子树
	root.Right = w
	return child
}

func rightRotation(root *TreapNode) *TreapNode {
	// 右旋， 左孩子上去
	child := root.Left
	// child的右孩子变更到父节点左子树下
	w := child.Right
	// 上去后， child的右孩子变成父节点
	child.Right = root
	// 父几点的左子树现在变成child的右子树
	root.Left = w
	return child
}

func Insert(root *TreapNode, x int) *TreapNode {
	if root == nil {
		return newNode(x)
	}

	if x > root.Val {
		root.Right = Insert(root.Right, x)
		// 如果优先级发生变化，则旋转
		if root.p > root.Right.p {
			root = leftRotation(root)
		}
	} else {
		root.Left = Insert(root.Left, x)
		// 如果优先级发生变化，则旋转
		if root.p > root.Left.p {
			root = rightRotation(root)
		}
	}

	return root
}

func Remove(root *TreapNode, x int) *TreapNode {
	if root == nil {
		return nil
	}

	if x < root.Val {
		root.Left = Remove(root.Left, x)
	} else if x > root.Val {
		root.Right = Remove(root.Right, x)
	} else if root.Left == nil {
		// 左子树为空， 右子树直接上位
		root = root.Right
	} else if root.Right == nil {
		// 右子树为空， 左子树直接上位
		root = root.Left
	} else if root.Left.p < root.Right.p {
		// 左右子树都不为空， 将优先级小的旋转上去，然后删除
		root = rightRotation(root)
		root.Right = Remove(root.Right, x)
	} else {
		root = leftRotation(root)
		root.Left = Remove(root.Left, x)
	}

	return root
}
