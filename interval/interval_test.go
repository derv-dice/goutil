package interval

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const div = "----------------------------------------------------------------------------------"

func TestInterval_Divide(t *testing.T) {
	now := time.Date(2022, 5, 7, 12, 47, 11, 0, time.UTC)
	interval := NewInterval(now, now.AddDate(5, 7, 0))

	is, err := interval.Divide(1, Year, time.Second)
	assert.NoError(t, err)
	fmt.Println(len(is))

	fmt.Printf("%s --> %s\n\n", now.Format(strLayout), now.Add(time.Hour*48).Format(strLayout))

	fmt.Println(div)
	for i := range is {
		x0, x1 := is[i].Vector()
		fmt.Println(is[i].String(), "|", x1.Sub(x0).String())
	}
	fmt.Println(div)
	fmt.Println()
}
