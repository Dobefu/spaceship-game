package game_scene

import (
	"github.com/Dobefu/spaceship-game/internal/player"
	"github.com/Dobefu/spaceship-game/internal/scene"
)

type GameScene struct {
	scene.Scene
}

func (s *GameScene) Init() {
	s.AddGameObject(player.NewPlayer(320, 320))
}
