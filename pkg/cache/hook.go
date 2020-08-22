package cache

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"

    "github.com/daodao97/egin/pkg/utils"
)

type logger struct {
}

func (l *logger) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
    newCtx := context.WithValue(ctx, "start", time.Now())
    return newCtx, nil
}

func (l *logger) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
    key := fmt.Sprintf("%s.%s", cmd.Args()[0], cmd.Args()[1])
    tc := time.Since(ctx.Value("start").(time.Time))
    utils.Logger.Channel("redis").Info(key, map[string]interface{}{
        "method": cmd.FullName(),
        "args":   cmd.Args(),
        "ums":    fmt.Sprintf("%v", tc),
    })
    return nil
}

func (l *logger) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
    fmt.Println("logger before pipe")
    return ctx, nil
}

func (l *logger) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
    fmt.Println("logger after pipe")
    return nil
}
