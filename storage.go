package main

import (
	"sync"
	"time"
)

type Storage struct {
	table  map[string]*node
	head   *node
	tail   *node
	size   int64
	mutex  *sync.Mutex
	config StorageConfig
}

func NewStorage(config StorageConfig) *Storage {
	return &Storage{
		table:  make(map[string]*node),
		head:   nil,
		tail:   nil,
		size:   0,
		mutex:  &sync.Mutex{},
		config: config,
	}
}

func (s *Storage) get(key string) ([]byte, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	n, ok := s.table[key]
	if !ok {
		return nil, false
	}

	if n.ttl.Before(time.Now()) {
		s.evict(n)
		s.size--
		return nil, false
	}

	return n.data, true
}

func (s *Storage) put(key string, data []byte) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	n, ok := s.table[key]

	ttl := time.Now().Add(s.config.Ttl)

	if ok {
		s.table[key].data = data
		s.table[key].ttl = ttl
		s.setHead(n)
		return true
	}

	for s.size >= s.config.Capacity {
		s.evict(s.tail)
		s.size--
	}

	newNode := &node{
		key:  key,
		data: data,
		ttl:  ttl,
	}

	s.table[key] = newNode
	s.setHead(newNode)
	s.size++

	return false
}

func (s *Storage) delete(key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	node, ok := s.table[key]
	if !ok {
		return false
	}

	s.evict(node)
	s.size--

	return true
}

func (s *Storage) clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for s.head != nil {
		s.evict(s.tail)
	}

	s.size = 0
}

func (s *Storage) Size() int64 {
	return s.size
}

func (s *Storage) Cap() int64 {
	return s.config.Capacity
}

func (s *Storage) RemoveExpired() int64 {
	var count int64 = 0
	for key, value := range s.table {
		if value.ttl.Before(time.Now()) {
			s.delete(key)
			count++
		}
	}
	return count
}

func (s *Storage) setHead(n *node) {
	if n == nil {
		return
	}

	if s.head == nil {
		s.head = n
		s.tail = n
		return
	}

	if s.head == n {
		return
	}

	if s.tail == n {
		s.tail = s.tail.prev
		s.tail.next = nil
		n.prev = nil
		s.head.prev = n
		n.next = s.head
		s.head = n
		return
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.prev = nil
	s.head.prev = n
	n.next = s.head
	s.head = n
}

func (s *Storage) evict(n *node) {
	if s.head == s.tail && s.head == n {
		s.head = nil
		s.tail = nil
		n.prev = nil
		n.next = nil
		delete(s.table, n.key)
		return
	}

	if s.head == n {
		s.head = s.head.next
		s.head.prev = nil
		n.prev = nil
		n.next = nil
		delete(s.table, n.key)
		return
	}

	if s.tail == n {
		s.tail = s.tail.prev
		s.tail.next = nil
		n.prev = nil
		n.next = nil
		delete(s.table, n.key)
		return
	}

	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.prev = nil
	n.next = nil
	delete(s.table, n.key)
}
