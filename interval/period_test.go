package interval

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPeriodMonth(t *testing.T) {
	now := time.Now()
	p := NewPeriodMonth(13, now)
	fmt.Println(p.Duration().Hours() / 24)
}

func TestNewPeriodYear(t *testing.T) {
	start := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	p := NewPeriodYear(10, start)
	fmt.Println(p.Duration().Hours() / 24)
	assert.Equal(t, 3653, int(p.Duration().Hours()/24))
}
