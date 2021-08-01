package binary_tree

import "banbo/CS/data_structure/binary_tree"

func findEQ(tNode *binary_tree.TreeNode, v int) *binary_tree.TreeNode {
	for tNode != nil {
		if v < tNode.Val {
			tNode = tNode.Left
		} else if v > tNode.Val {
			tNode = tNode.Right
		} else {
			return tNode
		}
	}
	return nil
}

// return closest value which great or equal than v.
func find(tNode *binary_tree.TreeNode, v int) int {
	var z *binary_tree.TreeNode = nil
	for tNode != nil {
		if v < tNode.Val {
			z = tNode
			tNode = tNode.Left
		} else if v > tNode.Val {
			tNode = tNode.Right
		} else {
			return v
		}
	}
	if z == nil {
		return -1
	}
	return z.Val
}

func findPre(root *binary_tree.TreeNode, v int) *binary_tree.TreeNode {
	var pre *binary_tree.TreeNode = nil
	for root != nil {
		pre = root
		if v < root.Val {
			root = root.Left
		} else if v > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return pre
}

func Add(root *binary_tree.TreeNode, v int) *binary_tree.TreeNode {
	if root == nil {
		return &binary_tree.TreeNode{Val: v}
	}
	pre := findPre(root, v)
	if pre.Val == v {
		return root
	}
	node := &binary_tree.TreeNode{Val: v}
	if node.Val < pre.Val {
		pre.Left = node
	} else {
		pre.Right = node
	}
	return root
}

func deleteMin(root *binary_tree.TreeNode) int {
	p := root
	for root.Left != nil {
		p = root
		root = root.Left
	}
	x := root.Val
	Remove(p, x)
	return x
}

func Remove(root *binary_tree.TreeNode, v int) *binary_tree.TreeNode {
	if root == nil {
		return nil
	}
	if v < root.Val {
		root.Left = Remove(root.Left, v)
	} else if v > root.Val {
		root.Right = Remove(root.Right, v)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			root.Val = deleteMin(root.Right)
		}
	}
	return root
}

func InOrder(root *binary_tree.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, ans)
	*ans = append(*ans, root.Val)
	InOrder(root.Right, ans)
}

func PreOrder(root *binary_tree.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	PreOrder(root.Left, ans)
	PreOrder(root.Right, ans)
}

func PostOrder(root *binary_tree.TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.Left, ans)
	PostOrder(root.Right, ans)
	*ans = append(*ans, root.Val)
}
