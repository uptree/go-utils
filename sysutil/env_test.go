package sysutil

import (
	"bytes"
	"os"
	"testing"

	"github.com/uptree/go-utils/assert"
)

func TestOS(t *testing.T) {
	if isw := IsWin(); isw {
		assert.Equal(t, true, isw)
		assert.Equal(t, true, IsWindows())
		assert.Equal(t, false, IsMac())
		assert.Equal(t, false, IsLinux())
	}
	if ism := IsMac(); ism {
		assert.Equal(t, true, ism)
		assert.Equal(t, false, IsWindows())
		assert.Equal(t, false, IsWin())
		assert.Equal(t, false, IsLinux())
	}
	if isl := IsLinux(); isl {
		assert.Equal(t, true, isl)
		assert.Equal(t, false, IsWindows())
		assert.Equal(t, false, IsWin())
		assert.Equal(t, false, IsMac())
	}
}

func TestConsole(t *testing.T) {
	assert.Equal(t, true, IsConsole(os.Stdout))
	assert.Equal(t, false, IsConsole(&bytes.Buffer{}))
}

func TestEnv(t *testing.T) {
	key := "EnvKey"
	value := "EnvValue"
	defValue := "DefValue"
	assert.Equal(t, nil, SetEnv(key, value))
	assert.Equal(t, value, GetEnv(key))
	assert.Equal(t, defValue, GetEnv("xx", defValue))
	for k, v := range EnvMap() {
		t.Logf("%s=%s", k, v)
	}
}
