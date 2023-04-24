package skiplist

import (
	"errors"
	"math/rand"
	"sync"
)

// TODO: generic type
type skipListNode struct {
	level int
	key   int
	value any
	// level -> next node
	next []*skipListNode
}

type SkipList struct {
	// max skip list currentMaxLevel
	maxLevel int
	// the probability of adding an additional layer to the currentMaxLevel
	p float64
	// pointer to the header node
	header *skipListNode
	// skiplist size
	size int
	// concurrency lock
	mu sync.Mutex
	// current currentMaxLevel for skiplist
	currentMaxLevel int
}

// create a new skiplist node
func (s *SkipList) newSkipNode(key int, value any) *skipListNode {
	node := &skipListNode{
		key:   key,
		value: value,
	}
	level := s.getRandomLevel()
	// create currentMaxLevel+1 slice for it
	node.next = make([]*skipListNode, level+1)
	return node
}

// get one new node's random currentMaxLevel
func (s *SkipList) getRandomLevel() int {
	point := int(s.p * 100)
	// default currentMaxLevel start from 1
	level := 1
	// randomly add the currentMaxLevel by the probability point
	for x := rand.Intn(100); x < point && level < s.maxLevel; {
		level++
	}
	return level
}

// Insert insert one element into skiplist
func (s *SkipList) Insert(key int, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// header is init

	// update is array which put node that the node.next[i] should be update to the new node
	update := make([]*skipListNode, s.maxLevel+1)

	current := s.header
	// start from the max currentMaxLevel of current skiplist
	for i := s.currentMaxLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		// find the last one who less than the `key` we want to insert
		update[i] = current
	}

}

// NewSkipList init a skiplist with maxLevel and probability
// probability should less than 1.00
func NewSkipList(opts ...SkiplistOption) (*SkipList, error) {
	option := &skiplistOption{
		maxLevel:    5,
		probability: 0.5,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.probability >= 1.0 {
		return nil, errors.New("invalid probability")
	}

	//init a start header node
	header := &skipListNode{
		key:   0,
		value: 0,
		// init start currentMaxLevel with nil pointer
		next:  make([]*skipListNode, option.maxLevel+1),
		level: option.maxLevel,
	}

	return &SkipList{
		maxLevel:        option.maxLevel,
		p:               option.probability,
		header:          header,
		size:            0,
		currentMaxLevel: 0,
	}, nil
}
