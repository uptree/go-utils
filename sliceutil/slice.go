package sliceutil

import (
	"strings"
)

// InSlice 判断字符串是否存在
func InSlice(s string, ss []string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// IsEmpty 判断Slice是否为空
func IsEmpty(ss []string) bool {
	if len(ss) == 0 {
		return true
	}
	return false
}

// Implode 别名 strings.Join
func Implode(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

// Explode 别名 strings.Split
func Explode(sep string, s string) []string {
	return strings.Split(s, sep)
}

// Unique slice去重
func Unique(ss []string) []string {
	ns := make([]string, 0)
	for _, s := range ss {
		if InSlice(s, ns) {
			break
		}
		ns = append(ns, s)
	}
	return ns
}

// Merge slice合并 - 不去重
func Merge(slice1, slice2 []string) []string {
	n := make([]string, len(slice1)+len(slice2))
	copy(n, slice1)
	copy(n[len(slice1):], slice2)
	return n
}

// Intersect slice交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// Union slice并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// Difference slice差集 - 属于slice1，不属于slice2
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	n := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			n = append(n, value)
		}
	}
	return n
}

// Filter slice过滤, 默认过滤空值
func Filter(ss []string, filter ...func(s string) bool) []string {
	var fn func(s string) bool
	if len(filter) > 0 && filter[0] != nil {
		fn = filter[0]
	} else {
		fn = func(s string) bool {
			return s != ""
		}
	}

	ns := make([]string, 0, len(ss))
	for _, s := range ss {
		if fn(s) {
			ns = append(ns, s)
		}
	}
	return ns
}

// IntersectUint64 slice交集
func IntersectUint64(slice1, slice2 []uint64) []uint64 {
	m := make(map[uint64]int)
	n := make([]uint64, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// UnionUint64 slice并集
func UnionUint64(slice1, slice2 []uint64) []uint64 {
	m := make(map[uint64]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// DifferenceUint64 slice差集 - 属于slice1，不属于slice2
func DifferenceUint64(slice1, slice2 []uint64) []uint64 {
	m := make(map[uint64]int)
	n := make([]uint64, 0)
	inter := IntersectUint64(slice1, slice2)
	for _, i := range inter {
		m[i]++
	}

	for _, v := range slice1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}
	return n
}

// IntersectInterface slice交集
func IntersectInterface(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	n := make([]interface{}, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// UnionInterface slice并集
func UnionInterface(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// DifferenceInterface slice差集 - 属于slice1，不属于slice2
func DifferenceInterface(slice1, slice2 []interface{}) []interface{} {
	m := make(map[interface{}]int)
	n := make([]interface{}, 0)
	inter := IntersectInterface(slice1, slice2)
	for _, i := range inter {
		m[i]++
	}

	for _, v := range slice1 {
		times, _ := m[v]
		if times == 0 {
			n = append(n, v)
		}
	}
	return n
}
