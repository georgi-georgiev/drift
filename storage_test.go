package main

import (
	"testing"
	"time"
)

func TestLRUEvictPolicy(t *testing.T) {
	ttl := time.Duration(time.Hour * 24)
	var capacity int64 = 10
	config := StorageConfig{ttl, capacity}
	cache := NewStorage(config)

	cache.put("key1", []byte("1jqoweijgn3120nvc0qjew0j"))
	cache.put("key2", []byte("2jqoweijgn3120nvc0qjew0j"))
	cache.put("key3", []byte("3jqoweijgn3120nvc0qjew0j"))
	cache.put("key4", []byte("4jqoweijgn3120nvc0qjew0j"))
	cache.put("key5", []byte("5jqoweijgn3120nvc0qjew0j"))
	cache.put("key6", []byte("6jqoweijgn3120nvc0qjew0j"))
	cache.put("key7", []byte("7jqoweijgn3120nvc0qjew0j"))
	cache.put("key8", []byte("8jqoweijgn3120nvc0qjew0j"))
	cache.put("key9", []byte("9jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))

	curSize := cache.Size()
	if curSize != 9 {
		t.Errorf("curSize should be 9, got %d\n", curSize)
	}

	cache.put("key10", []byte("10jqoweijgn3120nvc0qjew0j"))
	cache.put("key11", []byte("11jqoweijgn3120nvc0qjew0j"))

	curSize = cache.Size()
	if curSize != 10 {
		t.Errorf("curSize should be 10, got %d\n", curSize)
	}

	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("2102jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjw0j"))
	cache.put("key12", []byte("jqoweijgn3120nvc0qjgfejwqoijew0j"))
	cache.put("key13", []byte("jqoweijgnfqjweiojo120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120owiehqgvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn310nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn312oigjwqiej0nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))
	cache.put("key1", []byte("jqoweijgn3120nvc0qjew0j"))

	_, hit := cache.get("key2")
	if hit {
		t.Error("by lru policy, key2 should be evicted")
	}

	curSize = cache.Size()
	if curSize != 10 {
		t.Errorf("curSize should be 10, got %d\n", curSize)
	}

	_, hit = cache.get("key7")
	if !hit {
		t.Error("key 7 should not be evicted")
	}
}

func TestDeletion(t *testing.T) {
	ttl := time.Duration(time.Hour * 24)
	var capacity int64 = 10
	config := StorageConfig{ttl, capacity}
	cache := NewStorage(config)

	cache.put("key1", []byte("1jqoweijgn3120nvc0qjew0j"))
	cache.put("key2", []byte("2jqoweijgn3120nvc0qjew0j"))
	cache.put("key3", []byte("3jqoweijgn3120nvc0qjew0j"))
	cache.put("key4", []byte("4jqoweijgn3120nvc0qjew0j"))
	cache.put("key5", []byte("5jqoweijgn3120nvc0qjew0j"))
	cache.put("key6", []byte("6jqoweijgn3120nvc0qjew0j"))
	cache.put("key7", []byte("7jqoweijgn3120nvc0qjew0j"))
	cache.put("key8", []byte("8jqoweijgn3120nvc0qjew0j"))

	hit := cache.delete("key9")
	if hit {
		t.Error("never put key9, it sait hit key9")
	}

	hit = cache.delete("key8")
	if !hit {
		t.Error("key 8 should exist")
	}
	curSize := cache.Size()
	if curSize != 7 {
		t.Error("size doesn't right")
	}

	cache.clear()

	if cache.delete("key6") {
		t.Error("it didn't cleared")
	}
	curSize = cache.Size()
	if curSize != 0 {
		t.Error("size is not 0")
	}
}

func TestTTL(t *testing.T) {
	ttl := time.Duration(time.Second * 1)
	var capacity int64 = 10
	config := StorageConfig{ttl, capacity}
	cache := NewStorage(config)

	cache.put("key1", []byte("1jqoweijgn3120nvc0qjew0j"))
	cache.put("key2", []byte("2jqoweijgn3120nvc0qjew0j"))
	cache.put("key3", []byte("3jqoweijgn3120nvc0qjew0j"))
	cache.put("key4", []byte("4jqoweijgn3120nvc0qjew0j"))
	cache.put("key5", []byte("5jqoweijgn3120nvc0qjew0j"))
	cache.put("key6", []byte("6jqoweijgn3120nvc0qjew0j"))
	cache.put("key7", []byte("7jqoweijgn3120nvc0qjew0j"))
	cache.put("key8", []byte("8jqoweijgn3120nvc0qjew0j"))

	time.Sleep(time.Second * 2)

	_, hit := cache.get("key1")
	if hit {
		t.Error("Key 1 should be expired")
	}

	curSize := cache.Size()
	if curSize != 7 {
		t.Log(curSize)
		t.Error("even if expired, key will exist until it hits")
	}

	removeCount := cache.RemoveExpired()
	if removeCount != 7 {
		t.Error("after removeExpired, it will clear all")
	}
}
