package game

import (
	"net/http"
	"quick-start/pkg/nrandom"
	"quick-start/pkg/ntool"
	"sync"

	"github.com/SkyAPM/go2sky"
	xlogv2 "github.com/ivy-mobile/odin/xutil/xlog/v2"
	"github.com/olahol/melody"
)

type Game struct {
	nodeId              string                                //节点id
	ip                  string                                //节点ip
	logger              xlogv2.Logger                         //日志器
	diconnectHandler    func(uid int64, logger xlogv2.Logger) //玩家断开连接逻辑处理器，如需断开连接后处理业务，设置此函数即可
	afterRequestHandler func(ctx Context)                     //请求后处理逻辑处理器，如需请求后处理业务，设置此函数即可
	routes              sync.Map                              //消息处理器 key: version:route value: GameMessageHandler

	wsServer   *melody.Melody      //ws服务器
	httpServer *http.Server        //http服务器
	sessions   sync.Map            //会话管理
	urm        *UserRequestManager //用户请求管理

	// locator     locate.Locator          //定位器
	// transciever transciever.Transciever //传输器
	tracer *go2sky.Tracer //链路追踪器
}

func NewGame() *Game {
	IP, err := ntool.GetLocalIP()
	if err != nil {
		panic(err)
	}
	nodeId := nrandom.GetRandomString(6)
	return &Game{
		nodeId:   nodeId,
		ip:       IP,
		wsServer: melody.New(),
	}
}

func (g *Game) SetDisconnectHandler(handler func(uid int64, logger xlogv2.Logger)) {
	g.diconnectHandler = handler
}

func (g *Game) SetAfterRequestHandler(handler func(ctx Context)) {
	g.afterRequestHandler = handler
}

func (g *Game) GetNodeId() string {
	return g.nodeId
}

func (g *Game) GetIp() string {
	return g.ip
}
