package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileLineByLine(t *testing.T) {
	rows, err := ReadFileLineByLine("./testdata/test.txt", false)
	assert.NoError(t, err)
	assert.Equal(t, 12, len(rows))

	rows, err = ReadFileLineByLine("./testdata/test.txt", true)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(rows))

	fmt.Println(rows)
}
