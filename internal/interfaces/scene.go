package interfaces

type Scene interface {
	Init()
	GetGameObjects() []GameObject
	GetPauseScreenGameObjects() []GameObject
	AddGameObject(gameObject GameObject)
	SetCamera(camera Camera)
	GetCamera() (camera Camera)
	InvalidateDepthOrder()
	SetIsPaused(isPaused bool)
	GetIsPaused() (isPaused bool)
	SetCanPause(canPause bool)
	GetCanPause() (canPause bool)
}
