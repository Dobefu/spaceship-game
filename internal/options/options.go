package options

var (
	GlobalOptions = Options{
		Height: 640,
		Width:  640,

		OutlineWidth: 1.5,
	}
)

type Options struct {
	Height int
	Width  int

	OutlineWidth float32
}
