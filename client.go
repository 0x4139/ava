package ava

import (
	"net"
	"time"
)
import "gopkg.in/redis.v5"

type FailoverOptions struct {
	// The master name.
	MasterName         string
	// A seed list of host:port addresses of sentinel nodes.
	SentinelAddrs      []string

	Password           string
	DB                 int

	MaxRetries         int

	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration

	PoolSize           int
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}
type Options struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network            string
	// host:port address.
	Addr               string

	// Dialer creates new network connection and has priority over
	// Network and Addr options.
	Dialer             func() (net.Conn, error)

	// An optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password           string
	// A database to be selected after connecting to server.
	DB                 int

	// The maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries         int

	// Sets the deadline for establishing new connections. If reached,
	// dial will fail with a timeout.
	// Default is 5 seconds.
	DialTimeout        time.Duration
	// Sets the deadline for socket reads. If reached, commands will
	// fail with a timeout instead of blocking.
	ReadTimeout        time.Duration
	// Sets the deadline for socket writes. If reached, commands will
	// fail with a timeout instead of blocking.
	WriteTimeout       time.Duration

	// The maximum number of socket connections.
	// Default is 10 connections.
	PoolSize           int
	// Specifies amount of time client waits for connection if all
	// connections are busy before returning an error.
	// Default is 1 second.
	PoolTimeout        time.Duration
	// Specifies amount of time after which client closes idle
	// connections. Should be less than server's timeout.
	// Default is to not close idle connections.
	IdleTimeout        time.Duration
	// The frequency of idle checks.
	// Default is 1 minute.
	IdleCheckFrequency time.Duration

	// Enables read queries for a connection to a Redis Cluster slave node.
	ReadOnly           bool
}

func newClient(options Options) *redis.Client {
	redisOptions:= redis.Options(options)
	return redis.NewClient(&redisOptions)
}
func newFailOverClient(options FailoverOptions) *redis.Client {
	redisFailoverOptions:=redis.FailoverOptions(options)
	return redis.NewFailoverClient(&redisFailoverOptions)
}
