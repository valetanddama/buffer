package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCapacityFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	assert.Equal(t, 10, fifoBuffer.Cap())

	fifoBuffer.Append("test1")
	assert.Equal(t, 10, fifoBuffer.Cap())

	for i := 0; i < 20; i++ {
		fifoBuffer.Append("test")
	}
	assert.Equal(t, 10, fifoBuffer.Cap())
}

func TestLengthFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	assert.Equal(t, 0, fifoBuffer.Len())

	fifoBuffer.Append("test1")
	assert.Equal(t, 1, fifoBuffer.Len())

	for i := 0; i < 20; i++ {
		fifoBuffer.Append("test")
		assert.LessOrEqual(t, fifoBuffer.Len(), 10)
	}
	assert.Equal(t, 10, fifoBuffer.Len())
}

func TestAppendElementInFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	fifoBuffer.Append("test1")
	fifoBuffer.Append("test2")

	assert.Equal(t, []interface{}{"test1", "test2"}, fifoBuffer.GetItems())
}

func TestGetLastElementFromFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	fifoBuffer.Append("test1")
	fifoBuffer.Append("test2")
	fifoBuffer.Append("test3")
	fifoBuffer.Append("test4")

	assert.Equal(t, "test4", fifoBuffer.Last())
}

func TestShiftElementFromFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	fifoBuffer.Append("test1")
	fifoBuffer.Append("test2")
	fifoBuffer.Append("test3")
	fifoBuffer.Append("test4")

	fifoBuffer.Shift()
	assert.Equal(t, []interface{}{"test2", "test3", "test4"}, fifoBuffer.GetItems())
	assert.Equal(t, 3, fifoBuffer.Len())
	assert.Equal(t, 10, fifoBuffer.Cap())
}

func TestIfFullFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	fifoBuffer.Append("test1")
	fifoBuffer.Append("test2")
	fifoBuffer.Append("test3")
	fifoBuffer.Append("test4")

	assert.Equal(t, false, fifoBuffer.Full())

	for i := 0; i < 20; i++ {
		fifoBuffer.Append("test")
	}
	assert.Equal(t, true, fifoBuffer.Full())
}

func TestIfEmptyFifoBuffer(t *testing.T) {
	fifoBuffer := NewFifoBuffer(10)
	assert.Equal(t, true, fifoBuffer.Empty())

	fifoBuffer.Append("test")
	assert.Equal(t, false, fifoBuffer.Empty())
}

func BenchmarkAppendElementInFifoBuffer(b *testing.B) {
	fifoBuffer := NewFifoBuffer(100)

	for i := 0; i < b.N; i++ {
		fifoBuffer.Append("test1")
	}
}

func BenchmarkAppendElementInFullFifoBuffer(b *testing.B) {
	fifoBuffer := NewFifoBuffer(100)

	for i := 0; i < 100; i++ {
		fifoBuffer.Append("test")
	}

	for i := 0; i < b.N; i++ {
		fifoBuffer.Append("test1")
	}
}

func BenchmarkGetLastElementFromFifoBuffer(b *testing.B) {
	fifoBuffer := NewFifoBuffer(100)
	fifoBuffer.Append("test1")
	fifoBuffer.Append("test2")
	fifoBuffer.Append("test3")

	for i := 0; i < b.N; i++ {
		fifoBuffer.Last()
	}
}
