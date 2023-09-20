package maputil

import (
	"strings"

	"github.com/uptree/go-utils/convert"
)

// IsEmpty ...
func IsEmpty(mp map[string]string) bool {
	return len(mp) == 0
}

// HasKey ...
func HasKey(mp map[string]string, key string) bool {
	if _, ok := mp[key]; ok {
		return true
	}
	return false
}

// Value ...
func Value(mp map[string]string, key string) string {
	if HasKey(mp, key) {
		return mp[key]
	}
	return ""
}

// HasValue ...
func HasValue(mp map[string]string, value string) bool {
	for _, v := range mp {
		if v == value {
			return true
		}
	}
	return false
}

// Keys ...
func Keys(mp map[string]string) []string {
	ks := make([]string, 0, len(mp))
	for k, _ := range mp {
		ks = append(ks, k)
	}
	return ks
}

// Values ...
func Values(mp map[string]string) []string {
	vs := make([]string, 0, len(mp))
	for _, v := range mp {
		vs = append(vs, v)
	}
	return vs
}

// KeyToLower convert keys to lower case.
func KeyToLower(mp map[string]string) map[string]string {
	nmp := make(map[string]string, len(mp))
	for k, v := range mp {
		k = strings.ToLower(k)
		nmp[k] = v
	}
	return nmp
}

// Merge src 会覆盖 dst
func Merge(src, dst map[string]string) map[string]string {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// ToStringMap convert map[string]interface{} to map[string]string
func ToStringMap(src map[string]interface{}) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = convert.AnyToString(v)
	}
	return dst
}
