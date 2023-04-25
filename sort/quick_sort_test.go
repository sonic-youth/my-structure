package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

type intSlice struct {
	data []int
}

func (s *intSlice) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *intSlice) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

func (s *intSlice) Len() int {
	return len(s.data)
}
func TestQuickSort(t *testing.T) {
	n := 100000
	rand.Seed(time.Now().Unix())
	seq := make([]int, 0, n)
	for i := 0; i < n; i++ {
		seq = append(seq, rand.Int())
	}

	data := &intSlice{
		data: seq,
	}
	QuickSort(data)
	for i := 0; i < data.Len()-1; i++ {
		assert.False(t, data.data[i] > data.data[i+1])
	}
}

func TestQuickSort2(t *testing.T) {
	origin := []int{123, 456, 789, 123, 457, 890, 54367}
	expected := []int{123, 123, 456, 457, 789, 890, 54367}

	data := &intSlice{
		data: origin,
	}
	QuickSort(data)
	for i := 0; i < data.Len(); i++ {
		assert.True(t, data.data[i] == expected[i])
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		n := 1000000
		rand.Seed(time.Now().Unix())
		seq := make([]int, 0, n)
		for i := 0; i < n; i++ {
			seq = append(seq, rand.Int())
		}
		b.StartTimer()
		QuickSort(&intSlice{data: seq})
	}
}
