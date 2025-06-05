package redislock

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func runScript(ctx context.Context, rdb redis.Scripter, script *redis.Script, keys []string, args ...interface{}) *redis.Cmd {
	r := script.EvalSha(ctx, rdb, keys, args...)
	if redis.HasErrorPrefix(r.Err(), "NOSCRIPT") {
		if err := script.Load(ctx, rdb).Err(); err != nil {
			r = script.Eval(ctx, rdb, keys, args...) // fallback to EVAL
		} else {
			r = script.EvalSha(ctx, rdb, keys, args...) // retry EVALSHA
		}
	}
	return r
}
