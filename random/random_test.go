package random

import (
	"testing"
)

func TestGetRandomSring(t *testing.T) {
	t.Logf("GetRandomString: %s", GetRandomString(4, "abide123456"))
	t.Logf("GetRandomString: %s", GetRandomString(6, "abcdefghijklmnopqrstuvwxyz0123456789"))
}

func TestGetRandomChars(t *testing.T) {
	t.Logf("GetRandomChars: %s", GetRandomChars(4))
	t.Logf("GetRandomChars: %s", GetRandomChars(6))
}

func TestGetRandomNumbers(t *testing.T) {
	t.Logf("GetRandomNumbers: %s", GetRandomNumbers(4))
	t.Logf("GetRandomNumbers: %s", GetRandomNumbers(6))
}

func TestGetRandomInt(t *testing.T) {
	t.Logf("GetRandomInt: %d", GetRandomInt(10))
	t.Logf("GetRandomInt: %d", GetRandomInt(10000))
}

func TestGetRandomSlice(t *testing.T) {

	s := GetRandomSlice(2, []interface{}{"水", "蜜", "🍑"})
	t.Logf("GetRandomSlice: %s", s)
}

func TestGetRandomInVisible(t *testing.T) {
	s := GetRandomInVisible(1)
	t.Logf("GetRandomInVisible: %d", len(s))
}
