package main

import (
	"fmt"
	"strings"
)

const limitOffsetTmpl = "%s limit %d offset %d;"
const zeroOffsetTmpl = "%s limit %d;"

func QueryWithLimitAndOffset(query string, limit, offset int) (res string) {
	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")

	if offset == 0 {
		res = fmt.Sprintf(zeroOffsetTmpl, query, limit)
		return
	}

	res = fmt.Sprintf(limitOffsetTmpl, query, limit, offset)

	return
}
