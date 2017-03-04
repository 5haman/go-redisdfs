package matcher

import (
	redis "gopkg.in/redis.v3"
)

type RedisDirectedGraph struct {
	client *redis.Client
}

func (self *RedisDirectedGraph) GetConnected(from string) []string {
	items, err := self.client.SMembers(from).Result()
	if err == redis.Nil {
		// doesn't exist
	} else if err != nil {
		panic(err)
	} else {
		return items
	}
	return make([]string, 0)
}

func (self *RedisDirectedGraph) AddEdge(from string, to string) {
	err := self.client.SAdd(from, to).Err()
	if err != nil {
		panic(err)
	}
}

func (self *RedisDirectedGraph) RemoveEdge(from string, to string) {
	err := self.client.SRem(from, to).Err()
	if err != nil {
		panic(err)
	}
}

func (self *RedisDirectedGraph) Connect(url string) (err error) {
	self.client = redis.NewClient(&redis.Options{
		Addr:     url,
		DB:       0,
	})

	// get caller check for error
	_, err = self.client.Ping().Result()
	return err
}