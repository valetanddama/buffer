package buffer

import (
	"container/list"
	"sync"
)

type FifoBuffer struct {
	sync.RWMutex
	items *list.List
	cap   int
}

func NewFifoBuffer(cap int) *FifoBuffer {
	return &FifoBuffer{items: list.New(), cap: cap}
}

func (s *FifoBuffer) First() interface{} {
	s.RLock()
	defer s.RUnlock()

	if s.Empty() {
		return nil
	}
	return s.items.Front().Value
}

func (s *FifoBuffer) Last() interface{} {
	s.RLock()
	defer s.RUnlock()

	if s.Empty() {
		return nil
	}
	return s.items.Back().Value
}

func (s *FifoBuffer) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.Full() {
		s.Unlock()
		s.Shift()
		s.Lock()
	}

	s.items.PushBack(item)
}

func (s *FifoBuffer) Shift() {
	s.Lock()
	defer s.Unlock()

	if s.Empty() {
		return
	}

	if firstContainerItem := s.items.Front(); firstContainerItem != nil {
		s.items.Remove(firstContainerItem)
	}
}

func (s *FifoBuffer) Full() bool {
	return s.Len() == s.cap
}

func (s *FifoBuffer) Empty() bool {
	return s.Len() == 0
}

func (s *FifoBuffer) GetItems() []interface{} {
	s.RLock()
	defer s.RUnlock()

	items := make([]interface{}, 0, s.items.Len())

	for item := s.items.Front(); item != nil; item = item.Next() {
		items = append(items, item.Value)
	}

	return items
}

func (s *FifoBuffer) Len() int {
	return s.items.Len()
}
