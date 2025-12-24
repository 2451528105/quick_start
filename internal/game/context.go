package game

import (
	"quick-start/internal/config"
	"quick-start/internal/pm"
	"quick-start/internal/sdk/route"

	xlogv2 "github.com/ivy-mobile/odin/xutil/xlog/v2"
	"github.com/olahol/melody"
)

type Context interface {
	Seq() uint64           //请求序列号
	Uid() int64            //玩家id
	Route() string         //请求路由
	GameId() int32         //游戏id
	MsgId() string         //消息id
	TimeStamp() int64      //时间戳
	Version() string       //版本号
	Logger() xlogv2.Logger //日志器
	// OKResp(string ...goproto.Message)
	// ErrorResp(consts.ErrCode, ...string)
	// Broadcast(tag string, data goproto.Message, uids ...int64)
	// TracingContext() context.Context
}

type requestContext struct {
	g       *Game
	logger  xlogv2.Logger
	session *melody.Session
	req     *pm.Request
}

func newRequestContext(g *Game, session *melody.Session, req *pm.Request) *requestContext {
	return &requestContext{
		g:       g,
		logger:  g.logger,
		session: session,
		req:     req,
	}
}

func (c *requestContext) validate() bool {
	req := c.req
	if req == nil {
		return false
	}
	if req.GetSeq() == 0 {
		return false
	}
	if req.GetUid() == 0 {
		return false
	}
	if req.GetRoute() == "" {
		return false
	}
	if req.GetGameId() != int32(config.Cfg.GameId) {
		return false
	}
	if req.GetMsgId() == "" {
		return false
	}
	if req.GetTimestamp() == 0 {
		return false
	}
	if req.GetVersion() != route.Version {
		return false
	}
	if req.GetRoute() == "" {
		return false
	}
	return true
}

func (c *requestContext) Seq() uint64 {
	return c.req.GetSeq()
}
func (c *requestContext) Uid() int64 {
	return c.req.GetUid()
}
func (c *requestContext) Route() string {
	return c.req.GetRoute()
}
func (c *requestContext) GameId() int32 {
	return c.req.GetGameId()
}
func (c *requestContext) MsgId() string {
	return c.req.GetMsgId()
}
func (c *requestContext) TimeStamp() int64 {
	return c.req.GetTimestamp()
}
func (c *requestContext) Version() string {
	return c.req.GetVersion()
}
func (c *requestContext) Logger() xlogv2.Logger {
	return c.logger
}
