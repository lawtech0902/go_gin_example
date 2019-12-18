package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/lawtech0902/go_gin_example/pkg/setting"
	"time"
)

// redis工具包，封装一些函数

var RedisConn *redis.Pool

// 初始化实例
func Setup() error {
	RedisConn = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			
			if setting.RedisSetting.Password != "" {
				if _, err = conn.Do("AUTH", setting.RedisSetting.Password); err != nil {
					_ = conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
	}
	
	return nil
}

// 设置kv pair
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	
	return nil
}

// 检查key是否存在
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	
	return exists
}

// 获取key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	
	return reply, nil
}

// 删除key
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	
	return redis.Bool(conn.Do("DEL", key))
}

// 批处理删除
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	
	return nil
}
