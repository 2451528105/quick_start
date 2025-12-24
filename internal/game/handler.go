package game

import (
	"fmt"
	"quick-start/internal/pm"

	"github.com/ivy-mobile/odin/encoding/proto"
	"github.com/olahol/melody"
)

type (
	GameMessageHandler func(g *Game, s *melody.Session, msg *pm.Request) error //游戏消息处理器
)

func Handler[I any](fn func(Context, *I)) GameMessageHandler {
	return func(g *Game, s *melody.Session, msg *pm.Request) error {
		var content I
		if len(msg.GetBody()) > 0 {
			if err := proto.Unmarshal(msg.GetBody(), &content); err != nil {
				return fmt.Errorf("unmarshal body failed, err: %v", err)
			}
		}
		if ctx := newRequestContext(g, s, msg); ctx.validate() {
			fn(ctx, &content)
			if g.afterRequestHandler != nil {
				g.afterRequestHandler(ctx)
			}
		}
		return nil
	}
}
