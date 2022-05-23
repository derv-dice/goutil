package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestScsStruct struct {
	Col1 int       `csv:"col_1"`
	Col2 string    `csv:"col_2"`
	Col3 time.Time `csv:"col_3"`
}

func TestWriterCSV(t *testing.T) {
	wr, err := NewWriterCSV("testdata/test.csv", TestScsStruct{}, true)
	if err != nil {
		t.FailNow()
	}

	err = wr.WriteRow(TestScsStruct{
		Col1: 123,
		Col2: "1111",
		Col3: time.Now(),
	})

	assert.NoError(t, err)
	defer wr.Close()
}
