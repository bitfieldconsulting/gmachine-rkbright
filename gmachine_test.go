package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", gotMemSize, wantMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHalt(t *testing.T) {
	g := gmachine.New()
	g.Run()
	if g.P != 1 {
		t.Errorf("want P == 1, got %d", g.P)
	}
	g.Run()
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestNop(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpNOP
	g.Run()
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestIncA(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpINCA
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestDecA(t *testing.T) {
	g := gmachine.New()
	g.A = 2
	g.Memory[0] = gmachine.OpDECA
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestSub3From2(t *testing.T) {
	g := gmachine.New()
	copy(g.Memory, []uint64{
		gmachine.OpINCA,
		gmachine.OpINCA,
		gmachine.OpINCA,
		gmachine.OpDECA,
		gmachine.OpDECA,
	})
	g.Run()
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}
