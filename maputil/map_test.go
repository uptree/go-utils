package maputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equalf(t, true, IsEmpty(nil), "FAIL")
	assert.Equalf(t, true, IsEmpty(map[string]string{}), "FAIL")
	assert.Equalf(t, false, IsEmpty(map[string]string{"": ""}), "FAIL")
}

func TestHasKey(t *testing.T) {
	assert.Equalf(t, true, HasKey(map[string]string{"k": "v"}, "k"), "FAIL")
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
