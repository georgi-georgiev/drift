package main

import "time"

type node struct {
	key  string
	data []byte
	ttl  time.Time
	prev *node
	next *node
}
