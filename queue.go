package ava

import "gopkg.in/redis.v5"

type Queue struct {
	client *redis.Client
}

func NewQueue(options Options) *Queue {
	return &Queue{
		client:newClient(options),
	}
}

func NewFailOverQueue(options FailoverOptions) *Queue {
	return &Queue{
		client:newFailOverClient(options),
	}
}