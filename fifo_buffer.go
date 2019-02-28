package fifo_buffer

import "sync"

type FifoStack struct {
	sync.RWMutex
	items []interface{}
}

func NewFifoStack(cap int) *FifoStack {
	return &FifoStack{items: make([]interface{}, 0, cap)}
}

func (s *FifoStack) Last() interface{} {
	s.RLock()
	defer s.RUnlock()

	if s.Len() == 0 {
		return nil
	}
	return s.items[s.Len()-1]
}

func (s *FifoStack) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()

	if s.Len() == s.Cap() {
		s.Unlock()
		s.Shift()
		s.Lock()
	}
	s.items = append(s.items, item)
}

func (s *FifoStack) Shift() {
	s.Lock()
	defer s.Unlock()

	if s.Empty() {
		return
	}

	items := make([]interface{}, s.Len()-1, s.Cap())
	copy(items, s.items[1:])

	s.items = items
}

func (s *FifoStack) Full() bool {
	return s.Len() == s.Cap()
}

func (s *FifoStack) Empty() bool {
	return s.Len() == 0
}

func (s *FifoStack) GetItems() []interface{} {
	return s.items
}

func (s *FifoStack) Len() int {
	return len(s.items)
}

func (s *FifoStack) Cap() int {
	return cap(s.items)
}
