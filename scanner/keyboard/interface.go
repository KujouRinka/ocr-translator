package keyboard

const (
	Alt = 1 << iota
	Ctrl
	Shift
	Win
)

type Key struct {
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
}

func NewKey(modifiers int, keyCode int) Key {
	return Key{Modifiers: modifiers, KeyCode: keyCode}
}

type HotKeyListener interface {
	ListenOnKey(key Key) (<-chan struct{}, error)
	Run() error
}
