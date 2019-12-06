package redix

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	Address  string `yaml:"address" json:"address"`
	Port     int    `yaml:"port" json:"port"`
	Database int    `yaml:"database" json:"database"`

	connection redis.Conn
}

func (r *Redis) Connect() error {
	dial := fmt.Sprintf("%s:%v", r.Address, r.Port)
	conn, err := redis.Dial("tcp", dial, redis.DialDatabase(r.Database))
	if err != nil {
		return err
	}
	r.connection = conn
	return nil
}

func (r *Redis) Disconnect() error {
	if err := r.connection.Close(); err != nil {
		return err
	}
	r.connection = nil
	return nil
}

func (r Redis) Exists(key string) (bool, error) {
	reply, err := r.connection.Do("Exists", key)
	if err != nil {
		return false, err
	}
	return reply.(int64) == int64(1), nil
}

func (r Redis) Set(key string, jsonValue []byte) (interface{}, error) {
	return r.connection.Do("SET", key, jsonValue)
}

func (r Redis) HSet(key, field string, jsonValue []byte) (interface{}, error) {
	return r.connection.Do("HSET", key, field, jsonValue)
}

func (r Redis) Get(key string) (out string, err error) {
	reply, err := r.connection.Do("GET", key)
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", nil
	} else {
		return string(reply.([]byte)), nil
	}
}

func (r Redis) HGet(key, field string) (out string, err error) {
	reply, err := r.connection.Do("HGET", key, field)
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", nil
	} else {
		return string(reply.([]byte)), nil
	}
}
