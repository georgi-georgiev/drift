package main

import "time"

type StorageConfig struct {
	Ttl      time.Duration
	Capacity int64
}

type RedisConfig struct {
	RedisAddress  string
	RedisPassword string
	Channel       string
	Ttl           time.Duration
	Capacity      int64
}
