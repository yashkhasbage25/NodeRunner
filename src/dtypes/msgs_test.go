package dtypes

import (
	"testing"
)

func TestGetStr(t *testing.T) {
	posSamples := []struct {
		Pos    Position
		Answer string
	}{
		{Pos: Position{X: 20, Y: 50}, Answer: " X:20 Y:50"},
		{Pos: Position{X: 0, Y: 0}, Answer: " X:0 Y:0"},
		{Pos: Position{X: 845, Y: 600}, Answer: " X:845 Y:600"},
		{Pos: Position{X: -10, Y: -10}, Answer: " X:-10 Y:-10"},
	}
	for _, sample := range posSamples {
		got := sample.Pos.GetStr()
		want := sample.Answer
		if got != want {
			t.Error("(*Position) GetStr was incorrect, got:", got, "want:", want)
		}
	}

	evSamples := []struct {
		ev     *Event
		answer string
	}{
		{ev: &Event{
			EventType: "Update",
			Object:    "p1",
			P1Pos: Position{
				X: 900,
				Y: 420,
			},
			P2Pos: Position{
				X: 1000,
				Y: 420,
			},
			B1Pos: Position{
				X: 0,
				Y: 24,
			},
			B2Pos: Position{
				X: 0,
				Y: 24,
			},
			B3Pos: Position{
				X: 0,
				Y: 24,
			},
			G1Pos: Position{
				X: 215,
				Y: 305,
			},
			G2Pos: Position{
				X: 515,
				Y: 305,
			},
			G3Pos: Position{
				X: 465,
				Y: 305,
			},
			G4Pos: Position{
				X: 615,
				Y: 420,
			},
			P1Health: 20,
			P2Health: 50,
		},
			answer: " EventType:Update Object:p1 P1Pos: X:900 Y:420 P2Pos: X:1000 Y:420 B1pos: X:0 Y:24 B2Pos: X:0 Y:24 B3Pos: X:0 Y:24 G1Pos: X:215 Y:305 G2Pos: X:515 Y:305 G3Pos: X:465 Y:305 G4Pos: X:615 Y:420 P1Health:20 P2Health:50",
		},
	}
	for _, sample := range evSamples {
		got := sample.ev.GetStr()
		want := sample.answer
		if got != want {
			t.Error("(*Event) GetStr was incorrect, want: len",len(want), want, " \ngot: len", len(got), got)
		}
	}
}
