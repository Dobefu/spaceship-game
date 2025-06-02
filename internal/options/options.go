package options

var (
	GlobalOptions = Options{
		Height: 640,
		Width:  640,
	}
)

type Options struct {
	Height int
	Width  int
}
