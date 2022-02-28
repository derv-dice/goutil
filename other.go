package main

import (
	"math"
)

type LimitedSliceWindow struct {
	chunk         int
	from          int
	to            int
	arrayOfBounds [][2]int
}

func (w *LimitedSliceWindow) Bounds() [][2]int {
	return w.arrayOfBounds
}

func NewLimitedSliceWindow(chunk, from, to int) *LimitedSliceWindow {
	if chunk < 0 || from < 0 || to < 0 || (to < from) {
		return nil
	}

	res := &LimitedSliceWindow{
		chunk:         chunk,
		from:          from,
		to:            to,
		arrayOfBounds: [][2]int{},
	}

	allLen := to - from

	f := from
	t := f + chunk
	for i := 0; i < ChunksCount(allLen, chunk); i++ {
		if t >= allLen {
			t = allLen
		}

		res.arrayOfBounds = append(res.arrayOfBounds, [2]int{f, t})

		f = t
		t += chunk
	}

	return res
}

func ChunksCount(arrayLen, chunkLen int) int {
	n := float64(arrayLen) / float64(chunkLen)
	if math.Mod(n, 1) > 0 {
		return int(n) + 1
	}
	return int(n)
}
