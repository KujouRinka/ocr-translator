package keyboard

import "testing"

func TestKeyboardWindows(t *testing.T) {
	listener, err := NewListener()
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	ch1, err := listener.ListenOnKey(NewKey(Ctrl+Alt, 'M'))
	if err != nil {
		t.Fatal(err)
	}
	ch2, err := listener.ListenOnKey(NewKey(Ctrl+Alt, 'O'))
	if err != nil {
		t.Fatal(err)
	}

	go listener.Run()

	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			t.Log("Ctrl+Alt+M pressed")
		case <-ch2:
			t.Log("Ctrl+Alt+O pressed")
		}
	}
}
