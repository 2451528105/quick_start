package internal

import (
	"quick-start/internal/config"
	"quick-start/internal/game"
	"quick-start/internal/handle"
	"quick-start/internal/sdk/route"
)

func Init() {
	// 1. 初始化配置
	config.Init()
	// 2. 初始化游戏
	g := game.NewGame()
	// 3. 设置断开连接处理器
	g.SetDisconnectHandler(handle.DisconnectHandler)
	// 4. 设置请求后处理器
	g.SetAfterRequestHandler(handle.AfterRequestHandler)
	// 5. 注册消息处理器
	RegisterHandlers(route.Version, g)
	// 6. 启动游戏
	g.Start()
}

func RegisterHandlers(version string, g *game.Game) {
	g.RegisterHandler(version, route.Route_LoginGame, game.Handler(handle.LoginGame))
}
