package keyboard

import (
	"fmt"
	"syscall"
	"unsafe"
)

// ref: https://stackoverflow.com/questions/38646794/implement-a-global-hotkey-in-golang

type MSG struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

type Listener struct {
	id     int
	idChan map[int]chan struct{}

	user32           *syscall.DLL
	regHotKeyHandler *syscall.Proc
	getMsgHandler    *syscall.Proc
}

func NewListener() (*Listener, error) {
	user32, err := syscall.LoadDLL("user32")
	if err != nil {
		return nil, err
	}
	regHotKeyHandler, err := user32.FindProc("RegisterHotKey")
	if err != nil {
		return nil, err
	}
	getMsgHandler, err := user32.FindProc("GetMessageW")
	if err != nil {
		return nil, err
	}

	return &Listener{
		id:     1,
		idChan: make(map[int]chan struct{}),

		user32:           user32,
		regHotKeyHandler: regHotKeyHandler,
		getMsgHandler:    getMsgHandler,
	}, nil
}

func (l *Listener) ListenOnKey(key Key) (<-chan struct{}, error) {
	id := l.newId()
	err := l.regHotKey(id, key.Modifiers, key.KeyCode)
	if err != nil {
		return nil, err
	}

	ch := make(chan struct{})
	l.idChan[id] = ch

	return ch, nil
}

func (l *Listener) Run() error {
	for {
		msg, err := l.getMsg()
		// fmt.Printf("msg: %v\n", msg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		id := msg.WPARAM
		if id == 0 {
			continue
		}

		ch, ok := l.idChan[int(id)]
		if !ok {
			continue
		}

		ch <- struct{}{}
	}
}

func (l *Listener) Close() error {
	return l.user32.Release()
}

func (l *Listener) newId() int {
	ret := l.id
	l.id++
	return ret
}

func (l *Listener) regHotKey(id, modifiers, keyCode int) error {
	r1, _, err := l.regHotKeyHandler.Call(
		0,
		uintptr(id),
		uintptr(modifiers),
		uintptr(keyCode),
	)
	if r1 != 1 {
		return err
	}
	return nil
}

func (l *Listener) getMsg() (*MSG, error) {
	msg := &MSG{}
	l.getMsgHandler.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)
	// if r1 != 0 {
	// 	fmt.Println(r1)
	// 	return nil, err
	// }
	return msg, nil
}
