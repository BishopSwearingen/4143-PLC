package mascot_test

import (
	"testing"

	"github.com/BishopSwearingen/4143-PLC/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong Mascot :( ")
	}
}
