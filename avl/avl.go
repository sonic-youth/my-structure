package avl

type AVLTree struct {
	root *AVLNode
}

// TODO: add genericity type
type AVLNode struct {
	key   int
	value interface{}
	// height of the node
	height int
	// left and right child node
	lChild *AVLNode
	rChild *AVLNode
}

func (a *AVLTree) getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (a *AVLTree) abs(x int, y int) int {
	result := x - y
	if result < 0 {
		result = 0 - result
	}
	return result
}

func (a *AVLTree) max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// rotate right for this node
//
//	     y                               x
//	    / \                            /   \
//	   x  T4    rotate right         z      y
//	  / \       - - - - - - - ->    / \    / \
//	 z   T3                       T1  T2  T3 T4
//	/ \
//
// T1  T2
// input the node is y
// return which the root node(x) should point to
func (a *AVLTree) rotateR(y *AVLNode) *AVLNode {
	x := y.lChild
	t3 := x.rChild
	// set y as x's right child
	x.rChild = y
	// set t3 as y's left child
	y.lChild = t3
	// update height
	y.height = a.max(a.getHeight(y.lChild), a.getHeight(y.rChild)) + 1
	x.height = a.max(a.getHeight(x.lChild), y.height) + 1
	// return node x (the root node)
	return x
}

// rotate left for this node
//
//	      y                                   x
//	     / \                                /   \
//	    T4  x          rotate left         y     z
//	       / \       - - - - - - - ->     / \   / \
//	      T3  z                          T4 T3 T2 T1
//		     / \
//	        T2 T1
//
// input node is node y
// result is the root node(x) should point to
func (a *AVLTree) rotateL(y *AVLNode) *AVLNode {
	x := y.rChild
	t3 := x.lChild
	// set y as x's left child
	x.lChild = y
	// set t3 as y's right child
	y.rChild = t3
	// update height
	y.height = a.max(a.getHeight(y.lChild), a.getHeight(y.rChild)) + 1
	x.height = a.max(a.getHeight(x.rChild), y.height) + 1
	// return node x(the root node)
	return x
}

// rotate right-left for this node
/*
     y                      y                           z
    / \                    / \                        /   \
   T1  x   rotate right   T1  z    rotate left       y     x
      / \  - - - - - - >     / \   - - - - - - - >  / \   / \
     z  T2                  T3  x                  T1 T3 T4 T2
    / \                        / \
   T3  T4                     T4 T2
*/
func (a *AVLTree) rotateRL(y *AVLNode) *AVLNode {
	//1. rotate right for node x
	x := y.rChild
	z := a.rotateR(x)
	// set z as y's right child
	y.rChild = z
	//2. rotate left for node y
	z = a.rotateL(y)
	return z
}

// rotate left-right for this node
//
//	    y                       y                         z
//	   / \                     / \                      /   \
//	  x  T1    rotate left    z  T1   rotate right    x      y
//	 / \      - - - - - >    / \      - - - - - - >  / \    / \
//	T2  z                   x  T4                   T2 T3  T4 T1
//	   / \                 / \
//	  T3 T4               T2 T3
//
// input is the node y
// result is the root node (node z) point to
func (a *AVLTree) rotateLR(y *AVLNode) *AVLNode {
	//1. rotate left for node x
	x := y.lChild
	z := a.rotateL(x)
	// set z as y's left child
	y.lChild = z
	//2. rotate right for node y
	z = a.rotateR(y)
	return z
}

// insert one node into avl tree
// input: root is the root node of this tree/subtree, key/value
// return: return the new root node of this tree/subtree
func (a *AVLTree) insert(root *AVLNode, key int, value interface{}) *AVLNode {
	// if tree is empty, new a new node at root
	if root == nil {
		a.root = a.newAVLNode(key, nil, nil, value, 1)
		return a.root
	}
	if key < root.key {
		// insert to left tree
		root.lChild = a.insert(root.lChild, key, value)
	} else {
		// insert to right tree
		root.rChild = a.insert(root.rChild, key, value)
	}

	// finish insert, judge if the tree is balanced
	root.height = a.max(a.getHeight(root.lChild), a.getHeight(root.rChild)) + 1
	// if unbalanced, judge what rotate operation we need to do
	return a.balanceInsert(root, key)
}

func (a *AVLTree) Insert(key int, value interface{}) {
	a.root = a.insert(a.root, key, value)
}

// delete one node from avl tree
// input: root is the root node of this tree/subtree
// output: return the real root node when delete operation is done
func (a *AVLTree) delete(root *AVLNode, key int) *AVLNode {
	// if tree is empty, nothing to do
	if root == nil {
		return nil
	}
	// delete this root node
	// need to select a min value node at root's right subtree, and set it to root
	// then rotate to keep balance
	if key == root.key {
		// if there's no right subtree, set it as left child
		if root.rChild == nil {
			root = root.lChild
			return root
		} else {
			// select min value node at root's right subtree
			minValueNode := root.rChild
			for minValueNode.lChild != nil {
				minValueNode = minValueNode.lChild
			}
			// replace root with minValue Node
			root.key = minValueNode.key
			root.value = minValueNode.value
			// delete the minValue node (the minvalue node replaced the root node, so the real minValueNode need to delete)
			root.rChild = a.delete(root.rChild, minValueNode.key)
		}
	} else if key < root.key {
		root.lChild = a.delete(root.lChild, key)
	} else {
		root.rChild = a.delete(root.rChild, key)
	}

	// update height, judge whether the root is balanced
	root.height = a.max(a.getHeight(root.lChild), a.getHeight(root.rChild)) + 1
	return a.balanceDelete(root)
}

func (a *AVLTree) Delete(key int) {
	a.root = a.delete(a.root, key)
}

func (a *AVLTree) balanceInsert(root *AVLNode, key int) *AVLNode {
	// if insert to left
	if a.getHeight(root.lChild)-a.getHeight(root.rChild) == 2 {
		// if insert to left-left, rotate right
		if key < root.lChild.key {
			root = a.rotateR(root)
		} else {
			//if insert to left-right, rotate left-right
			root = a.rotateLR(root)
		}
	} else if a.getHeight(root.rChild)-a.getHeight(root.lChild) == 2 {
		// if insert to right-right, rotate left
		// Noted: must add equal case, if equal, you need rotate left for right-right case
		if key >= root.rChild.key {
			root = a.rotateL(root)
		} else {
			// if insert to right-left, rotate right-left
			root = a.rotateRL(root)
		}
	}
	return root
}

func (a *AVLTree) balanceDelete(root *AVLNode) *AVLNode {
	if a.getHeight(root.lChild)-a.getHeight(root.rChild) == 2 {
		//if the left-left subtree is higher than left-right, rotate right
		if a.getHeight(root.lChild.lChild) >= a.getHeight(root.lChild.rChild) {
			root = a.rotateR(root)
		} else {
			// rotate left-right
			root = a.rotateLR(root)
		}
	} else if a.getHeight(root.rChild)-a.getHeight(root.lChild) == 2 {
		// if the right-right subtree is higher than right-left, rotate left
		if a.getHeight(root.rChild.rChild) >= a.getHeight(root.rChild.lChild) {
			root = a.rotateL(root)
		} else {
			//rotate right-left
			root = a.rotateRL(root)
		}
	}
	return root
}

// traversal this tree to inorder sequence
func (a *AVLTree) inOrder(root *AVLNode) []int {
	var seq []int
	if root != nil {
		seq = a.inOrder(root.lChild)
		seq = append(seq, root.key)
		seq = append(seq, a.inOrder(root.rChild)...)
	}
	return seq
}

func (a *AVLTree) Inorder() []int {
	return a.inOrder(a.root)
}

// Load load the value from avl tree by the key
// TODO: concurrency
func (a *AVLTree) Load(key int) (interface{}, bool) {
	return a.load(a.root, key)
}

func (a *AVLTree) load(root *AVLNode, key int) (interface{}, bool) {
	if root == nil {
		return nil, false
	}
	if root.key == key {
		return root.value, true
	}
	if key < root.key {
		return a.load(root.lChild, key)
	}
	return a.load(root.rChild, key)
}

func (a *AVLTree) newAVLNode(key int, lChild *AVLNode, rChild *AVLNode, value interface{}, height int) *AVLNode {
	return &AVLNode{
		key:    key,
		value:  value,
		lChild: lChild,
		rChild: rChild,
		height: height,
	}
}

func (a *AVLTree) Root() *AVLNode {
	return a.root
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}
