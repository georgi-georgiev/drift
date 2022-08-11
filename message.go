package main

import "encoding/json"

type RedisMessage struct {
	Host   string
	Action string
	Key    string
	Data   []byte
}

func (msg *RedisMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

func (msg *RedisMessage) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}
	return nil
}
