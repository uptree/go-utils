package maputil

import (
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equalf(t, true, IsEmpty(nil), "FAIL")
	assert.Equalf(t, true, IsEmpty(map[string]string{}), "FAIL")
	assert.Equalf(t, false, IsEmpty(map[string]string{"": ""}), "FAIL")
}

func TestHasKey(t *testing.T) {
	assert.Equalf(t, true, HasKey(map[string]string{"k": "v"}, "k"), "FAIL")
}

func TestHasValue(t *testing.T) {
	assert.Equalf(t, true, HasValue(map[string]string{"k": "v"}, "v"), "FAIL")
	assert.Equalf(t, false, HasValue(map[string]string{"k": "v"}, "k"), "FAIL")
}

func TestValue(t *testing.T) {
	assert.Equalf(t, "v", Value(map[string]string{"k": "v"}, "k"), "FAIL")
	assert.Equalf(t, "", Value(map[string]string{"k": "v"}, "v"), "FAIL")
}

func TestKeys(t *testing.T) {
	assert.Equalf(t, []string{"k"}, Keys(map[string]string{"k": "v"}), "FAIL")
}

func TestValues(t *testing.T) {
	assert.Equalf(t, []string{"v"}, Values(map[string]string{"k": "v"}), "FAIL")
}

func TestKeyToLower(t *testing.T) {
	assert.Equalf(t, map[string]string{"k": "v"},
		KeyToLower(map[string]string{"K": "v"}), "FAIL")
}

func TestMerge(t *testing.T) {
	assert.Equalf(t, map[string]string{"k": "v", "K": "v"},
		Merge(map[string]string{"K": "v"}, map[string]string{"k": "v"}), "FAIL")
}

func TestToStringMap(t *testing.T) {
	assert.Equalf(t, map[string]string{"K": "1", "k": "2"},
		ToStringMap(map[string]interface{}{"K": 1, "k": "2"}), "FAIL")
}
