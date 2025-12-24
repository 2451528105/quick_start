package handle

import (
	"quick-start/internal/game"
	"quick-start/internal/pm"
)

func LoginGame(ctx game.Context, msg *pm.LoginGameRequest) {
	ctx.Logger().Info().Msgf("login game. gameId:%d,uid:%d.", ctx.GameId(), ctx.Uid())
}
