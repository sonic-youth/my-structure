package avl

// TODO: add genericity type
type AVLTree struct {
	key   int
	value interface{}
	// height of the node
	height int
	// left and right child node
	lChild *AVLTree
	rChild *AVLTree
}

func (a *AVLTree) getHeight() int {
	return a.height
}

func (a *AVLTree) getBalanceFactor() int {

}

// rotate right for this node
func (a *AVLTree) rotateR() {

}

// rotate left for this node
func (a *AVLTree) rotateL() {

}

// rotate right-left for this node
func (a *AVLTree) rotateRL() {

}

// rotate left-right for this node
func (a *AVLTree) rotateLR() {

}

// insert one node into avl tree
func (a *AVLTree) Insert(key int, value interface{}) {

}

// delete one node from avl tree
func (a *AVLTree) Delete(key int) {

}

// traversal this tree to inorder sequence
func (a *AVLTree) InOrder() {

}

func NewAVLNode(key int, lChild *AVLTree, value interface{}, rChild *AVLTree, height int) *AVLTree {
	node := &AVLTree{
		key:    key,
		value:  value,
		lChild: lChild,
		rChild: rChild,
		height: height,
	}
	return node
}
