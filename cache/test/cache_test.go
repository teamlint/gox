package test

import (
	"testing"
	"time"

	"github.com/teamlint/gox/cache"
)

var caches cache.Cache

type Model struct {
	Username   string
	Age        int
	IsApproved bool
	CreateAt   time.Time
	DeleteAt   *time.Time
}

func init() {
	opt := map[interface{}]interface{}{
		"Addr":              "127.0.0.1:6379",
		"Password":          "",
		"Database":          "",
		"Prefix":            "",
		"DefaultExpiration": "20m",
		"CleanupInterval":   "30m",
	}
	// caches = cache.New(cache.TypeMemory, opt)
	caches = cache.New(cache.TypeRedis, opt)
}
func TestAll(t *testing.T) {
	data := map[string]interface{}{
		"sring":   "foo 中文",
		"int":     int(123),
		"int64":   int64(123456),
		"float64": float64(12345.678),
		"boolF":   false,
		"boolT":   true,
		"struct": Model{
			Username:   "venjiang",
			Age:        36,
			IsApproved: false,
			CreateAt:   time.Now(),
		},
	}
	t.Log("==============================")
	for k, v := range data {
		t.Logf("set key %v\n", k)
		err := caches.Set(k, v, 0)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("set key %v success\n", k)
		}
	}
	t.Log("==============================")
	for k, _ := range data {
		val, err := caches.Get(k)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("get key %v=%v\n", k, val)
		}
	}
	t.Log("==============================")
	// set
	k := "struct4"
	model := Model{
		Username:   "read cache",
		Age:        38,
		IsApproved: true,
		CreateAt:   time.Now(),
	}
	caches.Set(k, &model, cache.DefaultExpiration)
	// read
	val := Model{}
	err := caches.Read(k, &val)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("read key %v=%v\n", k, val)
	}
}
func TestExp(t *testing.T) {
	k := "exp"
	v := "this is a expiration test"
	t.Logf("set key %v expiration %v\n", k, 2)
	err := caches.Set(k, v, cache.DefaultExpiration)
	// err := caches.Set(k, v, 3)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("set key %v success\n", k)
	}
	time.Sleep(2 * time.Second)
	val, err := caches.Get(k)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("get key %v=%v\n", k, val)
	}

}
func TestPatternDelete(t *testing.T) {
	caches.Set("k-1", "k-1", cache.DefaultExpiration)
	caches.Set("k-2", "k-2", cache.DefaultExpiration)
	caches.Set("k-3", "k-3", cache.DefaultExpiration)
	caches.Set("f-1", "f-1", cache.DefaultExpiration)
	caches.Set("f-1-1", "f-1-1", cache.DefaultExpiration)
	caches.Set("f-2", "f-2", cache.DefaultExpiration)
	var err error
	err = caches.PatternDelete("k-*")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("behind PatternDelete")

	var val interface{}

	if val, err = caches.Get("k-1"); err != nil {
		t.Error(err)
	}
	t.Log(val)
	if val, err = caches.Get("k-2"); err != nil {
		t.Error(err)
	}
	t.Log(val)
	if val, err = caches.Get("k-3"); err != nil {
		t.Error(err)
	}
	t.Log(val)
	if val, err = caches.Get("f-1"); err != nil {
		t.Error(err)
	}
	t.Log(val)
	t.Log(val)
	if val, err = caches.Get("f-1-1"); err != nil {
		t.Error(err)
	}
	t.Log(val)
}
