package redis

import (
	"encoding/json"

	"github.com/mishudark/eventhus"

	"github.com/go-redis/redis"
)

// Client redis
type Client struct {
	client *redis.Client
}

// NewClient returns a Client to acces to rabbitmq
func NewClient(addr string, password string, db int) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return &Client{
		client: client,
	}
}

// Publish a event
func (c *Client) Publish(event eventhus.Event, bucket, subset string) error {
	// bucket is the redis streams key

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	args := redis.XAddArgs{
		Stream: bucket,
		Values: map[string]interface{}{"data": body},
	}

	c.client.XAdd(&args)

	return err
}
