package storage

import (
	"fmt"
	"sync"

	"github.com/unheilbar/artforintrovert_entry_task/internal/entities"
)

type cache struct {
	mx       *sync.Mutex
	capacity uint
	users    map[string]*node
	queue    *linkedList
}

type node struct {
	key  string
	val  entities.User
	prev *node
	next *node
}

// linked list itself isn't thread safe,
type linkedList struct {
	tail *node
	head *node
	size uint
}

func (c *cache) Delete(ID string) {
	c.lock()
	defer c.unlock()
	rmNode, ok := c.users[ID]
	if !ok {
		fmt.Println("cache miss delete")
		return
	}

	c.queue.rmNode(rmNode)
	delete(c.users, ID)
}

func (c *cache) Get(ID string) (entities.User, bool) {
	c.lock()
	defer c.unlock()
	userNode, ok := c.users[ID]
	if !ok {
		fmt.Println("cache miss get")
		return entities.User{}, false
	}

	c.queue.moveToEnd(userNode)
	return userNode.val, true
}

func (c *cache) Set(user entities.User) {
	c.lock()
	defer c.unlock()
	userNode, ok := c.users[user.ID]
	if ok {
		userNode.val = user
		c.queue.moveToEnd(userNode)
		return
	}

	if c.capacity == c.queue.size {
		key := c.purge()
		delete(c.users, key)
	}

	userNode = &node{key: user.ID, val: user}
	c.users[user.ID] = userNode
	c.queue.addToEnd(userNode)
}

func (ll *linkedList) removeFirst() {
	ll.size--

	if ll.head == nil {
		return
	}

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	ll.head = ll.head.next
	ll.head.prev = nil
}

func (ll *linkedList) moveToEnd(n *node) {
	ll.rmNode(n)
	ll.addToEnd(n)
}

func (ll *linkedList) rmNode(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if ll.tail == n {
		ll.tail = n.prev
	}

	if ll.head == n {
		ll.head = n.next
	}

	ll.size--
}

func (ll *linkedList) addToEnd(n *node) {
	ll.size++

	if ll.head == nil {
		ll.head = n
		ll.tail = n
		return
	}

	n.prev = ll.tail
	ll.tail.next = n
	ll.tail = n
}

func (c *cache) purge() string {
	key := c.queue.head.key

	c.queue.removeFirst()
	return key
}

func (c *cache) lock() {
	c.mx.Lock()
}

func (c *cache) unlock() {
	c.mx.Unlock()
}
