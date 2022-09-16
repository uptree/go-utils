package sliceutil

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestInSlice(t *testing.T) {
	res := InSlice("a", []string{"c", "d", "e", "a"})
	assert.Equalf(t, true, res, "failed")
}

func TestIsEmpty(t *testing.T) {
	res := IsEmpty([]string{"c", "d", "e", "a"})
	assert.Equalf(t, false, res, "failed")
}

func TestImplode(t *testing.T) {
	res := Implode("|", []string{"a", "b", "c", "d"}...)
	assert.Equalf(t, "a|b|c|d", res, "failed")
}

func TestExplode(t *testing.T) {
	res := Explode("|", "a|b|c|d")
	assert.Equalf(t, []string{"a", "b", "c", "d"}, res, "failed")
}

func TestUnique(t *testing.T) {
	res := Unique([]string{"a", "c", "b", "c"})
	assert.Equalf(t, 3, len(res), "failed")
}

func TestMerge(t *testing.T) {
	res := Merge([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"a", "c", "b", "c", "d", "e", "a"}, res, "failed")
}

func TestIntersect(t *testing.T) {
	res := Intersect([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"c", "a"}, res, "failed")
}

func TestUnion(t *testing.T) {
	res := Union([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"a", "c", "b", "d", "e"}, res, "failed")
}

func TestDifference(t *testing.T) {
	res := Difference([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"b"}, res, "failed")
}

func TestIntersectUint64(t *testing.T) {
	res := IntersectUint64([]uint64{1, 2, 4}, []uint64{1, 2, 3, 4})
	assert.Equalf(t, []uint64{1, 2, 4}, res, "failed")
}

func TestUnionUint64(t *testing.T) {
	res := UnionUint64([]uint64{1, 2, 4}, []uint64{1, 6, 8})
	assert.Equalf(t, []uint64{1, 2, 4, 6, 8}, res, "failed")
}

func TestDifferenceUint64(t *testing.T) {
	res := DifferenceUint64([]uint64{1, 2, 4}, []uint64{1, 6, 8})
	assert.Equalf(t, []uint64{2, 4}, res, "failed")
}

func TestIntersectInterface(t *testing.T) {
	res := IntersectInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{1, 4}, res, "failed")
}

func TestUnionInterface(t *testing.T) {
	res := UnionInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{1, "2", 4, 2, 3}, res, "failed")
}

func TestDifferenceInterface(t *testing.T) {
	res := DifferenceInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{"2"}, res, "failed")
}
