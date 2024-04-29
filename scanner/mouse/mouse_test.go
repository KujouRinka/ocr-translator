package mouse

import (
	"testing"
	"time"
)

func TestMouse(t *testing.T) {
	for i := 0; i < 10; i++ {
		x, y := mouse()
		if x < 0 || y < 0 {
			t.Errorf("Mouse location is invalid: %d, %d", x, y)
		}
		t.Logf("Mouse location: %d, %d", x, y)
		time.Sleep(500 * time.Millisecond)
	}
}
