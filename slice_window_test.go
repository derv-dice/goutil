package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitedSliceWindow_Bounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	lsw := NewLimitedSliceWindow(7, 0, len(arr))
	assert.Equal(t, 4, len(lsw.Bounds()))
	assert.Equal(t, [2]int{0, 7}, lsw.Bounds()[0])
	assert.Equal(t, [2]int{21, 22}, lsw.Bounds()[3])
}

func TestChunksCount(t *testing.T) {
	assert.Equal(t, 4, ChunksCount(22, 7))
	assert.Equal(t, 0, ChunksCount(0, 7))
	assert.Equal(t, 2, ChunksCount(10, 5))
}
