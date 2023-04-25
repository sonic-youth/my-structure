package concurrent_hash_ring

import (
	"crypto/sha256"
	"sort"
	"strconv"
	"sync"
)

const defaultVirtualSpotsNum = 200

type ConcurrentHashRing interface {
	Add(values ...Item)
	Delete(values ...Item)
	Get(key Item) (Item, bool)
}

// Item defines the item stores in concurrent hash ring
type Item interface {
	String() string
}

type ConcurrentHashRingConfig struct {
	VirtualSpotsNum int
}

type concurrentHashRing struct {
	lock sync.Mutex
	//virtual spots number
	virtualSpotsNum int
	//real spots node
	spots []int
	//node value map, spots -> value
	nodeValueMap map[int]Item
}

// Add update all the values into concurrent hash ring
func (c *concurrentHashRing) Add(values ...Item) {
	for _, value := range values {
		c.add(value)
	}

	// when add done, sort all the spots
	c.sortSpots()
}

func (c *concurrentHashRing) add(value Item) {
	c.lock.Lock()
	defer c.lock.Unlock()
	//range the virtual spots node
	for i := 0; i < c.virtualSpotsNum; i++ {
		//calculate the spot for this value
		spot := c.calculateSpot(value.String() + ":" + strconv.Itoa(i))
		//if the node is a new real node, add it
		if _, ok := c.nodeValueMap[spot]; !ok {
			c.spots = append(c.spots, spot)
		}
		//update the node value map
		c.nodeValueMap[spot] = value
	}
}

func (c *concurrentHashRing) sortSpots() {
	c.lock.Lock()
	defer c.lock.Unlock()
	sort.Ints(c.spots)
}

// Delete delete the values in concurrent hash ring
func (c *concurrentHashRing) Delete(values ...Item) {
	for _, value := range values {
		c.delete(value)
	}
}

func (c *concurrentHashRing) delete(value Item) {
	c.lock.Lock()
	defer c.lock.Unlock()
	//TODO: optimize structure, no need to new a slice every time
	newSpots := make([]int, 0)
	//range the virtual spots node
	for _, spot := range c.spots {
		// if the spot is belong to the node who need to be delete
		// use string to compare
		if c.nodeValueMap[spot].String() == value.String() {
			delete(c.nodeValueMap, spot)
			continue
		}
		newSpots = append(newSpots, spot)
	}
	//just delete, no need to sort
	c.spots = newSpots
}

// Get will get the concurrent hash result item by given key
func (c *concurrentHashRing) Get(key Item) (Item, bool) {
	//calculate spot
	spot := c.calculateSpot(key.String())
	c.lock.Lock()
	defer c.lock.Unlock()

	//binary search
	idx := sort.Search(len(c.spots), func(i int) bool {
		return c.spots[i] >= spot
	})
	if idx == len(c.spots) {
		idx = 0
	}
	str, ok := c.nodeValueMap[c.spots[idx]]
	return str, ok
}

// calculateSpot calculate the spots of string value type
// TODO: change sha256 to a faster algorithm
func (c *concurrentHashRing) calculateSpot(value string) int {
	//calculate the spot for this conn
	hash := sha256.New()
	hash.Write([]byte(value))
	hashBytes := hash.Sum(nil)
	spot := 0
	for i := 0; i < sha256.Size; i += 4 {
		spot = spot ^ int((uint32(hashBytes[i])<<24)|(uint32(hashBytes[i+1])<<16)|(uint32(hashBytes[i+2])<<8)|(uint32(hashBytes[i+3])))
	}
	return spot
}

func NewConnHashRing(config *ConcurrentHashRingConfig) ConcurrentHashRing {
	if config.VirtualSpotsNum == 0 {
		config.VirtualSpotsNum = defaultVirtualSpotsNum
	}
	return &concurrentHashRing{
		virtualSpotsNum: config.VirtualSpotsNum,
		spots:           make([]int, 0),
		nodeValueMap:    make(map[int]Item),
	}
}
