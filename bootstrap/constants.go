package bootstrap

type Constants struct {
	Context Context
	// Add other constants as needed
}

type Context struct {
	Translator string
}

func NewConstants() *Constants {
	return &Constants{
		Context: Context{
			Translator: "translator",
		},
	}
}
