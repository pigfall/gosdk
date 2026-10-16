package utils

import (
	"sync"
)

// TailWriter buffers writes and retains only the last `capacity` bytes written.
type TailWriter struct {
	mu       sync.Mutex
	buf      []byte
	capacity int
	writeIdx int
	size     int
}

// NewTailWriter initializes a TailWriter with a fixed byte capacity.
func NewTailWriter(capacity int) *TailWriter {
	if capacity <= 0 {
		capacity = 1024 // Default capacity fallback
	}
	return &TailWriter{
		buf:      make([]byte, capacity),
		capacity: capacity,
	}
}

// Write implements the io.Writer interface.
func (tw *TailWriter) Write(p []byte) (n int, err error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	n = len(p)
	if n == 0 {
		return 0, nil
	}

	// If the write size exceeds capacity, keep only the end of p
	if n >= tw.capacity {
		p = p[n-tw.capacity:]
		copy(tw.buf, p)
		tw.writeIdx = 0
		tw.size = tw.capacity
		return n, nil
	}

	// Copy data into circular buffer
	for _, b := range p {
		tw.buf[tw.writeIdx] = b
		tw.writeIdx = (tw.writeIdx + 1) % tw.capacity
		if tw.size < tw.capacity {
			tw.size++
		}
	}

	return n, nil
}

// Bytes returns the retained tail bytes in chronological order.
func (tw *TailWriter) Bytes() []byte {
	tw.mu.Lock()
	defer tw.mu.Unlock()

	result := make([]byte, tw.size)
	if tw.size < tw.capacity {
		copy(result, tw.buf[:tw.size])
	} else {
		// Read from writeIdx to end, then from 0 to writeIdx
		n := copy(result, tw.buf[tw.writeIdx:])
		copy(result[n:], tw.buf[:tw.writeIdx])
	}

	return result
}

// String returns the retained tail as a string.
func (tw *TailWriter) String() string {
	return string(tw.Bytes())
}

// Reset clears the internal buffer.
func (tw *TailWriter) Reset() {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	tw.writeIdx = 0
	tw.size = 0
}
