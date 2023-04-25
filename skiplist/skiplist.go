package skiplist

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// TODO: generic type
type skipListNode struct {
	level int
	key   int
	value any
	// level -> next node
	next []*skipListNode
}

type Item struct {
	Key   int
	Value any
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

// Insert insert one element into skiplist, not allow insert the same key
// if the key is exist, return false
// else will insert successfully, return true
/*
                           +------------+
                           |  insert 50 |
                           +------------+
level 4     +-->1+                                                      100
                 |
                 |                      insert +----+
level 3         1+-------->10+---------------> | 50 |          70       100
                                               |    |
                                               |    |
level 2         1          10         30       | 50 |          70       100
                                               |    |
                                               |    |
level 1         1    4     10         30       | 50 |          70       100
                                               |    |
                                               |    |
level 0         1    4   9 10         30   40  | 50 |  60      70       100
                                               +----+
*/
func (s *SkipList) Insert(key int, value any) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// header is init

	// update is array which put node that the node.next[i] should be update to the new node
	// so `update` is the last nodes who less than `key` in level i
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

	// now current is the last nodes who less than `key` in level 0
	// current.next[0] is the place we desire to insert `key`
	current = current.next[0]
	// if the key is exist, return false
	if current != nil && current.key == key {
		return false
	}
	// if we reach the end of skipList, or the key is not exist, insert into `current`

	// generate a new skip list node
	newNode := s.newSkipNode(key, value)
	// if random level is greater than current level, you need to point the additional part of `update` to header
	if newNode.level > s.currentMaxLevel {
		for i := newNode.level; i > s.currentMaxLevel; i-- {
			update[i] = s.header
		}
		// update the current max level
		s.currentMaxLevel = newNode.level
	}
	// insert new node between `update` and `update.next[i]`
	for i := 0; i <= newNode.level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
	s.size++

	return true
}

// Delete delete one node from skipList
func (s *SkipList) Delete(key int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// generate `update` slice as insert
	update := make([]*skipListNode, s.maxLevel+1)

	current := s.header
	//start from highest level of skip list
	for i := s.currentMaxLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		update[i] = current
	}

	// now current is the last nodes who less than `key` in level 0
	// current.next[0] is the node what we desire to delete, so current.next[0] should equal to key
	current = current.next[0]
	// if not equal, just return
	if current == nil || (current != nil && current.key != key) {
		return
	}
	// found the key, delete it
	for i := 0; i <= s.currentMaxLevel; i++ {
		//if at level i, next node is not the current node, break loop
		if update[i].next[i] != current {
			break
		}
		update[i].next[i] = current.next[i]
		// remove levels which have no elements
		for s.currentMaxLevel > 0 && s.header.next[s.currentMaxLevel] == nil {
			s.currentMaxLevel--
		}
	}
	s.size--
	return
}

// Search search the key in skipList
func (s *SkipList) Search(key int) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.header
	for i := s.currentMaxLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
	}

	// now we found the level 0 of the last one who less than `key`
	current = current.next[0]

	if current != nil && current.key == key {
		return current.value, true
	}
	return nil, false
}

// Update update the value to the given key
func (s *SkipList) Update(key int, value any) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.header

	for i := s.currentMaxLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
	}

	// now we found the level 0 of the last one who less than `key`
	current = current.next[0]
	if current != nil && current.key == key {
		current.value = value
		return true
	}
	return false
}

// Range return the list of elements who is larger than key
func (s *SkipList) Range(key int) ([]*Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.header

	for i := s.currentMaxLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
	}

	current = current.next[0]

	if current == nil {
		return nil, false
	}

	result := make([]*Item, 0)
	for current != nil {
		result = append(result, &Item{Key: current.key, Value: current.value})
		current = current.next[0]
	}
	return result, true
}

func (s *SkipList) Size() int {
	return s.size
}

// NewSkipList init a skiplist with maxLevel and probability
// probability should less than 1.00
func NewSkipList(opts ...SkiplistOption) (*SkipList, error) {
	option := &skiplistOption{
		maxLevel:    18,
		probability: 0.5,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.probability >= 1.0 {
		return nil, errors.New("probability can't equal or larger than 1.0")
	}
	if option.maxLevel > 64 {
		return nil, errors.New("maxLevel can't larger than 64")
	}

	//init a start header node
	header := &skipListNode{
		key:   0,
		value: 0,
		// init start currentMaxLevel with nil pointer
		next:  make([]*skipListNode, option.maxLevel+1),
		level: option.maxLevel,
	}

	//set random seed
	rand.Seed(time.Now().Unix())
	return &SkipList{
		maxLevel:        option.maxLevel,
		p:               option.probability,
		header:          header,
		size:            0,
		currentMaxLevel: 0,
	}, nil
}
