package red_black_tree

func (rbt *RedBlackTree) flipLeft(u *Node) *Node {
	u.color, u.Right.color = u.Right.color, u.color
	return rbt.rotateLeft(u)
}

func (rbt *RedBlackTree) flipRight(u *Node) *Node {
	u.color, u.Left.color = u.Left.color, u.color
	return rbt.rotateRight(u)
}

func Search(root *Node, v int) *Node {
	if root == nil || root.Val == v {
		return root
	}

	if v < root.Val {
		return Search(root.Left, v)
	} else {
		return Search(root.Right, v)
	}
}

func (rbt *RedBlackTree) add(root *Node, x int) bool {
	node := &Node{Val: x, color: RED}
	pre := rbt.findPre(node)
	added := rbt.bstAdd(pre, node)
	if added {
		rbt.addFixUp(node)
		rbt.n++
	}
	return added
}

func (rbt *RedBlackTree) addFixUp(node *Node) {
	for node.color == RED {
		if node == rbt.root { // 如果是根节点，变黑色 done
			node.color = BLACK
			return
		}
		p := node.Parent
		/*
			如果左节点为黑色， 则违反左偏红黑树性质:
			如果某个节点其左孩子为黑色，则右孩子也为黑色
		*/
		if p.Left.color == BLACK {
			/*
					p(unknown)         swapColor(p, node)          p(unknown)
					 /       \         rotateLeft(node)   	    	/       \
				   s(b）    node(r)  ----------------------->    node(r)    T4
				   /  \      /  \          node = p               /  \
				  T1  T2    T3  T4         p = node.Parnet      s(b) T3
			*/
			rbt.flipLeft(p)
			node = p
			p = node.Parent
			// node still red and p still unknown.
		}
		/*
			如果父节点为红色，则会违反红黑树 no red-red edge 性质：
			两个红色节点不能相邻
		*/
		if p.color == BLACK {
			return // no red-red edge done.
		}
		// 此时p为红色， 由于p本来已存在，不能违反 no red-red 属性， 所以 g 此时必定为黑色
		// 此时node是红色，p是红色， g是黑色， node 和 p 违反 no red-red edge.
		g := p.Parent // grand parent of node， black
		// WARNING: 此时Right 有可能是p， 但P是红色
		if g.Right.color == BLACK {
			/*
			   表示p是g的左子树， 此时将p换颜色并提上去， 黑色p的左子树此时为红色的node， 右子树为红色的g
			   红色p原来的左子树是红色的node， 现在放到黑色的p下面不违反任何性质
			   红色p原来的右子树和p本身不违反任何，现在放到红色g下面，也不违反任何性质
			   黑色p换到黑色g的位置，原有的性质没有变化
			*/
			rbt.flipRight(g)
			return
		} else {
			/*
			   右左偏红黑树的性质可知，g的左子树必然为红色
			   表示p是g的右子树或者g的右子树为红色
			   现在将p和g的右子树换成黑色， p和node不再违反 no red-red,
			   但此时g变成了红色，有可能和其祖先违反 no red-red
			*/
			pushBlack(g)
			node = g
		}
	}
}
