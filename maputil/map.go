package maputil

import "strings"

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

// Keys ...
func Keys(mp map[string]string) []string {
	keys := make([]string, 0, len(mp))
	for k, _ := range mp {
		keys = append(keys, k)
	}
	return keys
}

// Values ...
func Values(mp map[string]string) []string {
	values := make([]string, 0, len(mp))
	for _, v := range mp {
		values = append(values, v)
	}
	return values
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
