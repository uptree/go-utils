package mathutil

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestMaxInt64(t *testing.T) {
	assert.Equalf(t, int64(10), MaxInt64([]int64{2, 4, 6, 8, 10}...), "FAIL")
	assert.Equalf(t, int64(10), MaxInt64([]int64{2, 4, -1, 8, 10}...), "FAIL")
}

func TestMinInt64(t *testing.T) {
	assert.Equalf(t, int64(0), MinInt64([]int64{2, 4, 6, 8, 10, 0}...), "FAIL")
	assert.Equalf(t, int64(-1), MinInt64([]int64{2, 4, -1, 8, 10}...), "FAIL")
}
