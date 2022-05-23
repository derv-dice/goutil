package webpb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess_Inc(t *testing.T) {
	p := NewProgressBar(0, 10)
	assert.Equal(t, 0, p.val)

	p.Inc()
	assert.Equal(t, 1, p.val)
	assert.Equal(t, true, p.Updated())
	assert.Equal(t, false, p.Updated())
}

func TestProcess_Add(t *testing.T) {
	p := NewProgressBar(0, 10)
	assert.Equal(t, 0, p.val)

	p.Add(0)
	assert.Equal(t, 0, p.val)

	p.Add(1)
	assert.Equal(t, 1, p.val)
	assert.Equal(t, true, p.Updated())
	assert.Equal(t, false, p.Updated())

	p.Add(20)
	assert.Equal(t, 10, p.val)
	assert.Equal(t, true, p.Updated())
	assert.Equal(t, false, p.Updated())
}
