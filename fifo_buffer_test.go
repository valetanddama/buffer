package fifo_buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapacityFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	assert.Equal(t, 10, fifoStack.Cap())

	fifoStack.Append("test1")
	assert.Equal(t, 10, fifoStack.Cap())

	for i := 0; i < 20; i++ {
		fifoStack.Append("test")
	}
	assert.Equal(t, 10, fifoStack.Cap())
}

func TestLengthFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	assert.Equal(t, 0, fifoStack.Len())

	fifoStack.Append("test1")
	assert.Equal(t, 1, fifoStack.Len())

	for i := 0; i < 20; i++ {
		fifoStack.Append("test")
		assert.LessOrEqual(t, fifoStack.Len(), 10)
	}
	assert.Equal(t, 10, fifoStack.Len())
}

func TestAppendElementInFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	fifoStack.Append("test1")
	fifoStack.Append("test2")

	assert.Equal(t, []interface{}{"test1", "test2"}, fifoStack.GetItems())
}

func TestGetLastElementFromFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	fifoStack.Append("test1")
	fifoStack.Append("test2")
	fifoStack.Append("test3")
	fifoStack.Append("test4")

	assert.Equal(t, "test4", fifoStack.Last())
}

func TestShiftElementFromFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	fifoStack.Append("test1")
	fifoStack.Append("test2")
	fifoStack.Append("test3")
	fifoStack.Append("test4")

	fifoStack.Shift()
	assert.Equal(t, []interface{}{"test2", "test3", "test4"}, fifoStack.GetItems())
	assert.Equal(t, 3, fifoStack.Len())
	assert.Equal(t, 10, fifoStack.Cap())
}

func TestIfFullFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	fifoStack.Append("test1")
	fifoStack.Append("test2")
	fifoStack.Append("test3")
	fifoStack.Append("test4")

	assert.Equal(t, false, fifoStack.Full())

	for i := 0; i < 20; i++ {
		fifoStack.Append("test")
	}
	assert.Equal(t, true, fifoStack.Full())
}

func TestIfEmptyFifoStack(t *testing.T) {
	fifoStack := NewFifoStack(10)
	assert.Equal(t, true, fifoStack.Empty())

	fifoStack.Append("test")
	assert.Equal(t, false, fifoStack.Empty())
}

func BenchmarkAppendElementInFifoStack(b *testing.B) {
	fifoStack := NewFifoStack(100)

	for i := 0; i < b.N; i++ {
		fifoStack.Append("test1")
	}
}

func BenchmarkAppendElementInFullFifoStack(b *testing.B) {
	fifoStack := NewFifoStack(100)

	for i := 0; i < 100; i++ {
		fifoStack.Append("test")
	}

	for i := 0; i < b.N; i++ {
		fifoStack.Append("test1")
	}
}

func BenchmarkGetLastElementFromFifoStack(b *testing.B) {
	fifoStack := NewFifoStack(100)
	fifoStack.Append("test1")
	fifoStack.Append("test2")
	fifoStack.Append("test3")

	for i := 0; i < b.N; i++ {
		fifoStack.Last()
	}
}
