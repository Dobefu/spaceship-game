package globals

import (
	"log/slog"

	"github.com/quasilyte/gdata/v2"
)

var (
	DataManager *gdata.Manager
	Options     AvailableOptions
)

type AvailableOptions struct {
	EnableShaders bool
}

func init() {
	var err error
	DataManager, err = gdata.Open(gdata.Config{
		AppName: "spaceship-game",
	})

	if err != nil {
		slog.Error(err.Error())
	}

	enableShaders, _ := DataManager.LoadObjectProp("settings", "enable-shaders")

	Options = AvailableOptions{
		EnableShaders: len(enableShaders) <= 0 || enableShaders[0] != 0,
	}
}
