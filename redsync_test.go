package redsync

import (
	"os"
	"testing"
	"time"

	redigolib "github.com/gomodule/redigo/redis"
	"github.com/ridwanadhip/redsync/redis"
	"github.com/ridwanadhip/redsync/redis/redigo"
	"github.com/stvp/tempredis"
)

var servers []*tempredis.Server

type testCase struct {
	poolCount int
	pools     []redis.Pool
}

func makeCases(poolCount int) map[string]*testCase {
	return map[string]*testCase{
		"redigo": {
			poolCount,
			newMockPoolsRedigo(poolCount),
		},
	}
}

// Maintain separate blocks of servers for each type of driver
const SERVER_POOLS = 4
const SERVER_POOL_SIZE = 8
const REDIGO_BLOCK = 0

func TestMain(m *testing.M) {
	for i := 0; i < SERVER_POOL_SIZE*SERVER_POOLS; i++ {
		server, err := tempredis.Start(tempredis.Config{})
		if err != nil {
			panic(err)
		}
		servers = append(servers, server)
	}
	result := m.Run()
	for _, server := range servers {
		_ = server.Term()
	}
	os.Exit(result)
}

func TestRedsync(t *testing.T) {
	for k, v := range makeCases(8) {
		t.Run(k, func(t *testing.T) {
			rs := New(v.pools...)

			mutex := rs.NewMutex("test-redsync")
			_ = mutex.Lock()

			assertAcquired(t, v.pools, mutex)
		})
	}
}

func newMockPoolsRedigo(n int) []redis.Pool {
	pools := make([]redis.Pool, n)

	offset := REDIGO_BLOCK * SERVER_POOL_SIZE

	for i := 0; i < n; i++ {
		server := servers[i+offset]
		pools[i] = redigo.NewPool(&redigolib.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redigolib.Conn, error) {
				return redigolib.Dial("unix", server.Socket())
			},
			TestOnBorrow: func(c redigolib.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		})
	}
	return pools
}
