package redigo

import "github.com/ridwanadhip/redsync/redis"

var _ redis.Conn = (*conn)(nil)

var _ redis.Pool = (*pool)(nil)
