package main

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type Rdb struct {
	storage *Storage
	client  *redis.Client
	config  RedisConfig
	ctx     context.Context
	channel string
	host    string
}

func NewRdb(redisConfig RedisConfig, storageConfig StorageConfig) *Rdb {

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.RedisAddress,
		Password: redisConfig.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	storage := NewStorage(storageConfig)

	rdb := &Rdb{
		storage: storage,
		config:  redisConfig,
		ctx:     ctx,
		client:  client,
		channel: redisConfig.Channel,
		host:    hostName,
	}

	go rdb.Receive()

	return rdb
}

func (rdb *Rdb) Receive() {
	topic := rdb.client.Subscribe(rdb.ctx, rdb.config.Channel)

	channel := topic.Channel()

	for msg := range channel {
		payload := &RedisMessage{}
		err := payload.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}

		if payload.Host != rdb.host {
			switch payload.Action {
			case "set":
				rdb.storage.put(payload.Key, payload.Data)
				break
			case "del":
				rdb.storage.delete(payload.Key)
				break
			case "clear":
				rdb.storage.clear()
				break
			default:
				panic(payload.Action + " is not valid")
			}
		}
	}
}

func (rdb *Rdb) Get(key string) ([]byte, bool) {
	data, hit := rdb.storage.get(key)

	if !hit {
		stringCmd := rdb.client.Get(rdb.ctx, key)
		if stringCmd == nil {
			return nil, false
		}

		if stringCmd.Err() != nil {
			panic(stringCmd.Err())
		}

		rdb.storage.put(key, []byte(stringCmd.Val()))

		return []byte(stringCmd.Val()), false
	}

	return data, hit
}

func (rdb *Rdb) Put(key string, data []byte) bool {
	hit := rdb.storage.put(key, data)
	statusCmd := rdb.client.Set(rdb.ctx, key, data, rdb.config.Ttl)
	if statusCmd != nil && statusCmd.Err() != nil {
		panic(statusCmd.Err())
	}
	msg := RedisMessage{Host: rdb.host, Action: "set", Key: key, Data: data}
	msgBytes, err := msg.MarshalBinary()
	if err != nil {
		panic(err)
	}

	intCmd := rdb.client.Publish(rdb.ctx, rdb.channel, msgBytes)
	if intCmd != nil && intCmd.Err() != nil {
		panic(intCmd.Err())
	}
	return hit
}

func (rdb *Rdb) Delete(key string) bool {
	hit := rdb.storage.delete(key)

	intCmd := rdb.client.Del(rdb.ctx, key)
	if intCmd != nil && intCmd.Err() != nil {
		panic(intCmd.Err())
	}

	msg := RedisMessage{Host: rdb.host, Action: "del", Key: key}
	msgBytes, err := msg.MarshalBinary()
	if err != nil {
		panic(err)
	}

	intCmd = rdb.client.Publish(rdb.ctx, rdb.channel, msgBytes)
	if intCmd != nil && intCmd.Err() != nil {
		panic(intCmd.Err())
	}

	return hit
}

func (rdb *Rdb) Clear() {
	rdb.storage.clear()
	rdb.client.FlushAll(rdb.ctx)
	msg := RedisMessage{Host: rdb.host, Action: "clear"}
	msgBytes, err := msg.MarshalBinary()
	if err != nil {
		panic(err)
	}
	intCmd := rdb.client.Publish(rdb.ctx, rdb.channel, msgBytes)
	if intCmd != nil && intCmd.Err() != nil {
		panic(intCmd.Err())
	}
}
