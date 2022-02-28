package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryWithLimitAndOffset(t *testing.T) {
	q := "select * from table_name where id = $1;"

	assert.Equal(t, "select * from table_name where id = $1 limit 100;",
		QueryWithLimitAndOffset(q, 100, 0),
	)

	assert.Equal(t, "select * from table_name where id = $1 limit 100 offset 200;",
		QueryWithLimitAndOffset(q, 100, 200),
	)
}
