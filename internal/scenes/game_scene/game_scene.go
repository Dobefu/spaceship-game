package game_scene

import (
	"github.com/Dobefu/spaceship-game/internal/player"
	"github.com/Dobefu/spaceship-game/internal/scene"
	"github.com/Dobefu/spaceship-game/internal/vectors"
)

type GameScene struct {
	scene.Scene
}

func (s *GameScene) Init() {
	s.AddGameObject(player.NewPlayer(vectors.Vector2{X: 320, Y: 320}))
}
