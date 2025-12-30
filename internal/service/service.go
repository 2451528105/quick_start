package service

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	xlogv2 "github.com/ivy-mobile/odin/xutil/xlog/v2"
	"github.com/redis/go-redis/v9"
)

type sendHandler interface {
	SendMessageByRMQ(uid int64, message []byte)
}

var SVC = &service{}

type service struct {
	ctx         context.Context  //上下文
	sendHandler sendHandler      //发送消息处理器
	logger      xlogv2.Logger    //日志器
	nodeId      string           //节点id
	redisClient *redis.Client    //redis客户端
	redisLocker *redsync.Redsync //分布式锁

}

func Init(ctx context.Context, redisClient *redis.Client, redisLocker *redsync.Redsync, nodeId string, logger xlogv2.Logger) {
	SVC.redisClient = redisClient
	SVC.redisLocker = redisLocker
	SVC.nodeId = nodeId
	SVC.logger = logger.With("module", "service")
	SVC.ctx = ctx
}
