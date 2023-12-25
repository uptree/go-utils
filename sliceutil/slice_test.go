package sliceutil

import (
	"fmt"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestInSlice(t *testing.T) {
	res := InSlice("a", []string{"c", "d", "e", "a"})
	assert.Equalf(t, true, res, "FAIL")
}

func TestIsEmpty(t *testing.T) {
	res := IsEmpty([]string{"c", "d", "e", "a"})
	assert.Equalf(t, false, res, "FAIL")
}

func TestImplode(t *testing.T) {
	res := Implode("|", []string{"a", "b", "c", "d"}...)
	assert.Equalf(t, "a|b|c|d", res, "FAIL")
}

func TestExplode(t *testing.T) {
	res := Explode("|", "a|b|c|d")
	assert.Equalf(t, []string{"a", "b", "c", "d"}, res, "FAIL")
}

func TestUnique(t *testing.T) {
	res := Unique([]string{"a", "a", "c", "b", "c", "c", "a"})
	fmt.Println(res)
	assert.Equalf(t, 3, len(res), "FAIL")
}

func TestMerge(t *testing.T) {
	res := Merge([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"a", "c", "b", "c", "d", "e", "a"}, res, "FAIL")
}

func TestIntersect(t *testing.T) {
	res := Intersect([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"c", "a"}, res, "FAIL")
}

func TestUnion(t *testing.T) {
	res := Union([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"a", "c", "b", "d", "e"}, res, "FAIL")
}

func TestDifference(t *testing.T) {
	res := Difference([]string{"a", "c", "b"}, []string{"c", "d", "e", "a"})
	assert.Equalf(t, []string{"b"}, res, "FAIL")
}

func TestFilter(t *testing.T) {
	// 默认过滤空值
	assert.Equalf(t, []string{"a", "c", "b"}, Filter([]string{"a", "c", "b", ""}), "FAIL")
	// 	自定义
	assert.Equalf(t, []string{"b"},
		Filter([]string{"a", "c", "b", ""},
			func(s string) bool {
				return s == "b"
			}),
		"FAIL")
}

func TestIntersectUint64(t *testing.T) {
	res := IntersectUint64([]uint64{1, 2, 4}, []uint64{1, 2, 3, 4})
	assert.Equalf(t, []uint64{1, 2, 4}, res, "FAIL")
}

func TestUnionUint64(t *testing.T) {
	res := UnionUint64([]uint64{1, 2, 4}, []uint64{1, 6, 8})
	assert.Equalf(t, []uint64{1, 2, 4, 6, 8}, res, "FAIL")
}

func TestDifferenceUint64(t *testing.T) {
	res := DifferenceUint64([]uint64{1, 2, 4}, []uint64{1, 6, 8})
	assert.Equalf(t, []uint64{2, 4}, res, "FAIL")
}

func TestIntersectInterface(t *testing.T) {
	res := IntersectInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{1, 4}, res, "FAIL")
}

func TestUnionInterface(t *testing.T) {
	res := UnionInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{1, "2", 4, 2, 3}, res, "FAIL")
}

func TestDifferenceInterface(t *testing.T) {
	res := DifferenceInterface([]interface{}{1, "2", 4}, []interface{}{1, 2, 3, 4})
	assert.Equalf(t, []interface{}{"2"}, res, "FAIL")
}

func TestUnshift(t *testing.T) {
	ss := []string{"a", "b", "c", "d"}
	length := Unshift("x", &ss)
	assert.Equalf(t, []string{"x", "a", "b", "c", "d"}, ss, "FAIL")
	assert.Equalf(t, 5, length, "FAIL")
}

func TestPush(t *testing.T) {
	ss := []string{"a", "b", "c", "d"}
	length := Push("x", &ss)
	assert.Equalf(t, []string{"a", "b", "c", "d", "x"}, ss, "FAIL")
	assert.Equalf(t, 5, length, "FAIL")
}

func TestShift(t *testing.T) {
	ss := []string{"a", "b", "c", "d"}
	res, ok := Shift(&ss)
	assert.Equalf(t, []string{"b", "c", "d"}, ss, "FAIL")
	assert.Equalf(t, "a", res, "FAIL")
	assert.Equalf(t, true, ok, "FAIL")
}

func TestPop(t *testing.T) {
	ss := []string{"a", "b", "c", "d"}
	res, ok := Pop(&ss)
	assert.Equalf(t, []string{"a", "b", "c"}, ss, "FAIL")
	assert.Equalf(t, "d", res, "FAIL")
	assert.Equalf(t, true, ok, "FAIL")
}

func TestSplitByLen(t *testing.T) {
	res := SplitByLen("a|b|c|d", 2)
	assert.Equalf(t, []string{"a|", "b|", "c|", "d"}, res, "FAIL")
	res = SplitByLen("钟山风雨起苍黄", 2)
	assert.Equalf(t, []string{}, res, "FAIL")
}
