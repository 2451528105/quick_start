package game

import (
	"context"
	"fmt"
	"quick-start/internal/config"
	"quick-start/internal/sdk/tools"
	"quick-start/internal/service"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"

	"github.com/ivy-mobile/odin/xutil/xfile"
	xlogv2 "github.com/ivy-mobile/odin/xutil/xlog/v2"
)

func (g *Game) Start() {
	var err error

	// 1.初始化日志
	logger := buildLogger(g.nodeId).With("env", config.Cfg.Env).With("node", g.nodeId).With("ip", g.ip)
	g.logger = logger.With("module", "game")

	// 2.初始化链路追踪
	if config.Cfg.SkyWalking.Enable {
		g.tracer, err = buildTracer(config.Cfg.SkyWalking.ServiceName, config.Cfg.SkyWalking.Addr)
		if err != nil {
			panic(err)
		}
	}

	// 3.初始化redis
	var redisClient *redis.Client
	redisClient, err = tools.InitRedis()
	if err != nil {
		panic(err)
	}

	// 4.初始化分布式锁
	redisLocker := redsync.New(goredis.NewPool(redisClient))

	// 5.初始化服务层
	service.Init(context.Background(), redisClient, redisLocker, g.nodeId, g.logger)

}

func (g *Game) RegisterHandler(version, route string, handler GameMessageHandler) {
	key := fmt.Sprintf("%s:%s", version, route)
	g.routes.Store(key, handler)
}

// buildLogger 构建日志器
func buildLogger(node string) xlogv2.Logger {
	logCfg := config.Cfg.Log
	opts := make([]xlogv2.Option, 0)
	if logCfg != nil {
		opts = append(opts,
			xlogv2.WithLevel(logCfg.Level),
			xlogv2.WithMode(logCfg.Mode),
		)
		if logCfg.File != nil {
			fileCfg := logCfg.File
			fileName := xfile.JoinFilename(fileCfg.FilePath, "-", node)
			opts = append(opts,
				xlogv2.WithFile(fileName, fileCfg.MaxSize, fileCfg.MaxBackups, fileCfg.MaxAge, fileCfg.Compress, fileCfg.LocalTime))
		}
	}
	return xlogv2.New(opts...)
}

// buildTracer 构建链路追踪器
func buildTracer(serviceName, oapAddr string) (*go2sky.Tracer, error) {
	// 创建reporter负责将追踪数据发送到OAP服务器
	r, err := reporter.NewGRPCReporter(oapAddr)
	if err != nil {
		return nil, err
	}

	// 创建tracer负责生成和传递追踪数据 同时设置了服务名
	tracer, err := go2sky.NewTracer(serviceName, go2sky.WithReporter(r))
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
