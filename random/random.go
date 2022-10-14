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

const (
	CharUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"         // 26
	CharLowerCase = "abcdefghijklmnopqrstuvwxyz"         // 26
	CharNumber    = "0123456789"                         // 10
	CharSymbol    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" // 32
)

// GetRandomInt 生成值小于max的随机数
func GetRandomInt(max int) int {
	rand.Seed(getRandSeek())
	return rand.Intn(max)
}

// GetRandomChars 生成英文字母随机字符串
func GetRandomChars(num int) string {
	return GetRandomString(num, CharUpperCase+CharLowerCase)
}

// GetRandomNumbers 生成数字类型随机字符串
func GetRandomNumbers(num int) string {
	return GetRandomString(num, CharNumber)
}

// GetRandomString 生成随机字符串
func GetRandomString(num int, ss ...string) string {
	s := CharNumber + CharLowerCase
	if len(ss) > 0 {
		s = ss[0]
	}
	le := len(s)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(le)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}

// GetRandomSlice 随机取num个slice元素，允许重复取
func GetRandomSlice(num int, ss []interface{}) []interface{} {
	ns := make([]interface{}, 0)
	r := rand.New(rand.NewSource(getRandSeek()))
	for i := 0; i < num; i++ {
		ns = append(ns, ss[r.Intn(len(ss))])
	}
	return ns
}

func GetRandomInVisible(num int) string {
	//Zs
	inVisibleChar := []rune{
		'\u2000',
		'\u2001',
		'\u2002',
		'\u2003',
		'\u2004',
		'\u2005',
		'\u2006',
		'\u2007',
		'\u2008',
		'\u2009',
		'\u200A',
		'\u202F',
		'\u205F',
		'\u3000',
	}
	cc := make([]rune, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		cc = append(cc, inVisibleChar[r.Intn(len(inVisibleChar))])
	}
	return string(cc)
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
