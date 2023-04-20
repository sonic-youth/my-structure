package avl

import (
	"github.com/smartystreets/goconvey/convey"
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
	rand.Seed(time.Now().Unix())
	convey.Convey("test avl tree", t, func() {
		convey.Convey("test for insert and delete", func() {
			n := 100000
			var seq []int
			for i := 0; i < n; i++ {
				seq = append(seq, rand.Int())
			}
			avl := NewAVLTree()
			for _, x := range seq {
				avl.Insert(x, x)
			}
			result := avl.Inorder()
			flag := true
			for i := 0; i < len(result)-1; i++ {
				if result[i] > result[i+1] {
					flag = false
					break
				}
			}
			convey.So(flag, convey.ShouldBeTrue)
			for i := 0; i < 10000; i++ {
				avl.Delete(seq[i])
			}
			flag = true
			result = avl.Inorder()
			for i := 0; i < len(result)-1; i++ {
				if result[i] > result[i+1] {
					flag = false
					break
				}
			}
			convey.So(flag, convey.ShouldBeTrue)
		})
		convey.Convey("test for load value", func() {
			n := 100000
			var seq []int
			for i := 0; i < n; i++ {
				seq = append(seq, rand.Int())
			}
			avl := NewAVLTree()
			for _, x := range seq {
				avl.Insert(x, x)
			}
			flag := true
			for i := 0; i < 1000; i++ {
				value, ok := avl.Load(seq[n-i-1])
				if !ok || value.(int) != seq[n-i-1] {
					flag = false
					break
				}
			}
			convey.So(flag, convey.ShouldBeTrue)
		})
	})
}

func BenchmarkAVLTree_Insert(b *testing.B) {
	avl := NewAVLTree()
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		avl.Insert(i, i)
	}
}

func BenchmarkAVLTree_Delete(b *testing.B) {
	avl := NewAVLTree()
	rand.Seed(time.Now().Unix())
	n := 1000000
	for i := 0; i < n; i++ {
		avl.Insert(rand.Intn(1000000), rand.Intn(1000000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Delete(rand.Intn(1000000))
	}
}

func BenchmarkAVLTree_Load(b *testing.B) {
	avl := NewAVLTree()
	rand.Seed(time.Now().Unix())
	n := 1000000
	for i := 0; i < n; i++ {
		avl.Insert(rand.Intn(1000000), rand.Intn(1000000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Load(rand.Intn(n))
	}
}
