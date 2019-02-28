package buffer

import "sync"

type FifoBuffer struct {
	sync.RWMutex
	items []interface{}
}

func NewFifoBuffer(cap int) *FifoBuffer {
	return &FifoBuffer{items: make([]interface{}, 0, cap)}
}

func (s *FifoBuffer) Last() interface{} {
	s.RLock()
	defer s.RUnlock()

	if s.Len() == 0 {
		return nil
	}
	return s.items[s.Len()-1]
}

func (s *FifoBuffer) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.Len() == s.Cap() {
		s.Unlock()
		s.Shift()
		s.Lock()
	}
	s.items = append(s.items, item)
}

func (s *FifoBuffer) Shift() {
	s.Lock()
	defer s.Unlock()

	if s.Empty() {
		return
	}

	items := make([]interface{}, s.Len()-1, s.Cap())
	copy(items, s.items[1:])

	s.items = items
}

func (s *FifoBuffer) Full() bool {
	return s.Len() == s.Cap()
}

func (s *FifoBuffer) Empty() bool {
	return s.Len() == 0
}

func (s *FifoBuffer) GetItems() []interface{} {
	return s.items
}

func (s *FifoBuffer) Len() int {
	return len(s.items)
}

func (s *FifoBuffer) Cap() int {
	return cap(s.items)
}
