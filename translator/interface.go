package translator

type Engine interface {
	Translate(text string) (string, error)
}
