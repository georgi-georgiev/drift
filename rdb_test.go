package main

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	redisAddress := "localhost:6379"
	redisPassword := ""
	channel := "test"
	ttl := time.Duration(time.Hour * 24)
	var capacity int64 = 10
	config := RedisConfig{redisAddress, redisPassword, channel, ttl, capacity}
	config2 := StorageConfig{ttl, capacity}
	rdb := NewRdb(config, config2)

	rdb.Put("key1", []byte("1jqoweijgn3120nvc0qjew0j"))
	_, hit := rdb.Get("key1")
	if !hit {
		t.Error("key1 not found")
	}
}
