package cache

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	r "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func TestLoad(t *testing.T) {
	testKey := "test"

	option, err := r.ParseURL(os.Getenv("KV_URL"))
	if err != nil {
		t.Errorf("parse redis url failed:%v", err)
	}
	conf := redis.RedisConf{
		Host:        option.Addr,
		Type:        "node",
		Pass:        option.Password,
		Tls:         true,
		NonBlock:    false,
		PingTimeout: time.Second,
	}
	rds := redis.MustNewRedis(conf)
	ctx := context.Background()
	err = rds.SetCtx(ctx, testKey, "1")
	if err != nil {
		t.Errorf("set :%+v", err)
		logc.Error(ctx, err)
	}
	v, err := rds.GetCtx(ctx, testKey)
	if err != nil {
		logc.Error(ctx, err)
	}
	fmt.Println("get key's value:", v)
}
