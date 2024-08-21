package structs

import (
	"context"
	"errors"
	"log"

	"github.com/redis/go-redis/v9"
)

type Options struct {
	Repository string
	Addr       string
	Password   string
	DB         int
}

func (o *Options) Validate() error {
	if o.Addr == "" {
		return errors.New("redis address is required")
	}

	if o.Password == "" {
		return errors.New("redis password is required")
	}

	if o.Repository == "" {
		return errors.New("redis repository name is required")
	}

	return nil
}

func (o *Options) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     o.Addr,
		Password: o.Password,
		DB:       o.DB,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			cn.ClientSetName(ctx, o.Repository)

			log.Println("redis is connected with name :", cn.ClientGetName(ctx))
			return nil
		},
		PoolSize: 1,
	}
}
