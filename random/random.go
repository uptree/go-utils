package random

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

var (
	randSeek = int64(1)
	l        sync.Mutex
)

// GetRandomInt 生成值小于max的随机数
func GetRandomInt(max int) int {
	rand.Seed(getRandSeek())
	return rand.Intn(max)
}

// GetRandomChars 生成英文字母随机字符串
func GetRandomChars(num int) string {
	ss := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GetRandomString(num, ss)
}

// GetRandomNumbers 生成数字类型随机字符串
func GetRandomNumbers(num int) string {
	return GetRandomString(num)
}

// GetRandomString 生成随机字符串
func GetRandomString(num int, str ...string) string {
	s := "0123456789"
	if len(str) > 0 {
		s = str[0]
	}
	l := len(s)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}
