package skiplist

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSkipList(t *testing.T) {
	rand.Seed(time.Now().Unix())
	convey.Convey("test for skipList", t, func() {
		convey.Convey("test insert and delete", func() {
			n := 100
			seq := make([]int, 0)
			for i := 0; i < n; i++ {
				seq = append(seq, rand.Intn(1000))
			}
			skipList, err := NewSkipList()
			convey.So(err, convey.ShouldBeNil)
			for _, x := range seq {
				skipList.Insert(x, x)
			}
			deleted := make(map[int]bool)
			for i := 0; i < 50; i++ {
				skipList.Delete(seq[i])
				deleted[seq[i]] = true
			}
			for i := 50; i < n; i++ {
				if ok := deleted[seq[i]]; !ok {
					value, ok := skipList.Search(seq[i])
					convey.So(ok, convey.ShouldBeTrue)
					convey.So(value, convey.ShouldEqual, seq[i])
				}
			}
		})

		convey.Convey("test range search", func() {
			seq := []int{123, 123, 14436, 78678, 43543, 78769, 535, 654765, 898, 797, 797, 4124, 4654, 523, 4326, 57, 56, 876, 976, 835, 45, 246, 3, 7}
			skipList, err := NewSkipList()
			convey.So(err, convey.ShouldBeNil)

			for _, x := range seq {
				skipList.Insert(x, x)
			}
			result, ok := skipList.Range(123)
			convey.So(ok, convey.ShouldBeTrue)
			for _, x := range result {
				t.Logf("get item: %v", x)
			}
		})
	})
}

func BenchmarkSkipList_Search(b *testing.B) {
	skipList, err := NewSkipList()
	assert.Nil(b, err)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10000; i++ {
		x := rand.Intn(100000000000)
		skipList.Insert(x, rand.Intn(100000000000))
	}
	b.Logf("item for skipList: %v", skipList.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		skipList.Search(rand.Intn(100000000000))
	}
}
