package globals

var (
	Options AvailableOptions
)

type AvailableOptions struct {
	EnableShaders bool
}

func init() {
	Options = AvailableOptions{
		EnableShaders: true,
	}
}
