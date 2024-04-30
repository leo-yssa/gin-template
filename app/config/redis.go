package config

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

func NewStore() (redis.Store, error) {
	return redis.NewStore(10, "tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), "", []byte(os.Getenv("REDIS_SECRET")))
}