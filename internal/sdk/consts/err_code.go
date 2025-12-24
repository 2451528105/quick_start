package consts

import "fmt"

type ErrCode int32

func (e ErrCode) String() string {
	if e == 0 {
		return "A000000"
	}
	return fmt.Sprintf("G%06d", e)
}

const (
	OK ErrCode = 0
)

const (
	InternalServerError ErrCode = iota + 100000
	BadRequest                  // 请求失败
	Unauthorized                // 未授权
	RequestRpcFail              // 远程调用错误
	TxCommitErr                 // 事务提交错误
	OptDatabaseErr              // 数据库操作错误
	RequestErr                  // 错误的请求数据
	OptRedisErr                 // Redis操作错误
	ProtobufErr                 // Protobuf错误
	JsonMarshalErr              // Json序列化错误
	InvalidUserId               // 无效的用户ID
	InvalidTableId              // 无效的牌桌ID
	InvalidGameId               // 无效的游戏id
	InvalidOpt                  // 无效的操作
)
