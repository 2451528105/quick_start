package game

import (
	"fmt"
	"quick-start/internal/config"

	"github.com/ivy-mobile/odin/xutil/xfile"
	xlogv2 "github.com/ivy-mobile/odin/xutil/xlog/v2"
)

func (g *Game) Start() {

	// 1.初始化日志
	logger := buildLogger(g.nodeId).With("env", config.Cfg.Env).With("node", g.nodeId).With("ip", g.ip)
	logger.Info().Any("config", config.Cfg).Msg("----------- server config!")
	g.logger = logger.With("module", "game")

}

func (g *Game) RegisterHandler(version, route string, handler GameMessageHandler) {
	key := fmt.Sprintf("%s:%s", version, route)
	g.routes.Store(key, handler)
}

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
