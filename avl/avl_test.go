package avl

import (
	"math/rand"
	"testing"
	"time"
)

func TestInitAVLTree(t *testing.T) {
	seq := []int{976, 2664, 8498, 5132, 4589, 763, 412, 8268, 6643, 8560, 5584, 952, 7928, 508, 5455, 2502, 7575, 4942, 5013, 1436}
	avl := NewAVLTree()
	for _, x := range seq {
		avl.Insert(x, x)
		t.Logf("after insert %v: root: %+v, root-left:%+v, root-right:%+v, %v", x, avl.root, avl.root.lChild, avl.root.rChild, avl.Inorder())
	}
	t.Logf("inorder seq: %v", avl.Inorder())
}

func TestAVLTree(t *testing.T) {
	n := 100000
	var seq []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		seq = append(seq, rand.Int())
	}
	avl := NewAVLTree()
	for _, x := range seq {
		avl.Insert(x, x)
	}
	result := avl.Inorder()
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			t.Logf("invalid inorder result: %v", result[i])
		}
	}

	for i := 0; i < 1000; i++ {
		avl.Delete(seq[i])
	}

	result = avl.Inorder()
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			t.Logf("invalid inorder result:%v", result[i])
		}
	}
}
