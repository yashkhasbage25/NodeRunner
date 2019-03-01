package handler

import (
	"testing"

	"github.com/IITH-SBJoshi/concurrency-3/src/coords"
	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
)

func TestAllignedWithLadder(t *testing.T) {
	coords.Initialize()
	var sample_test1 [4]dtypes.Rect
	var answer [4]bool
	sample_test1[0] = dtypes.Rect{400, 350, 430, 390}
	sample_test1[1] = dtypes.Rect{405, 350, 435, 390}
	sample_test1[2] = dtypes.Rect{395, 350, 425, 390}
	sample_test1[3] = dtypes.Rect{300, 350, 330, 390}
	answer[0] = true
	answer[1] = true
	answer[2] = true
	answer[3] = false
	for i := 0; i < 4; i++ {
		got := AllignedWithLadder(sample_test1[i])
		want := answer[i]
		if got != want {
			t.Error("AllignedWithLadder was incorrect, got: ", got, "want: ", want)
		}
	}
}
func TestGetPosition(t *testing.T) {
	var sample_test2 [4]dtypes.Rect
	var answer2 [4]dtypes.Position
	sample_test2[0] = dtypes.Rect{400, 350, 430, 390}
	sample_test2[1] = dtypes.Rect{0, 600, 1200, 800}
	sample_test2[2] = dtypes.Rect{200, 350, 300, 550}
	sample_test2[3] = dtypes.Rect{300, 350, 330, 390}
	answer2[0] = dtypes.Position{415, 370}
	answer2[1] = dtypes.Position{600, 700}
	answer2[2] = dtypes.Position{250, 450}
	answer2[3] = dtypes.Position{315, 370}
	for i := 0; i < 4; i++ {
		got := GetPosition(sample_test2[i])
		want := answer2[i]
		if got != want {
			t.Error("Get Position was incorrect, got: ", got, "want: ", want)
		}
	}
}
func TestGetBoundary(t *testing.T) {
	var answer3 [4]dtypes.Rect
	var sample_test3 [4]dtypes.Position
	sample_test3[0] = dtypes.Position{415, 370}
	sample_test3[1] = dtypes.Position{600, 700}
	sample_test3[2] = dtypes.Position{250, 450}
	sample_test3[3] = dtypes.Position{315, 370}
	answer3[0] = dtypes.Rect{400, 350, 430, 390}
	answer3[1] = dtypes.Rect{585, 680, 615, 720}
	answer3[2] = dtypes.Rect{235, 430, 265, 470}
	answer3[3] = dtypes.Rect{300, 350, 330, 390}

	for i := 0; i < 4; i++ {
		got := GetBoundary(sample_test3[i])
		want := answer3[i]
		if got != want {
			t.Error("AllignedWithLadder was incorrect, got: ", got, "want: ", want)
		}
	}
}
func TestFallsFromBlock(t *testing.T) {
	var answer4 [4]bool
	var sample_test4 [4]dtypes.Rect
	sample_test4[0] = dtypes.Rect{510, 410, 540, 450}
	sample_test4[1] = dtypes.Rect{536, 410, 566, 450}
	sample_test4[2] = dtypes.Rect{545, 410, 575, 450}
	sample_test4[3] = dtypes.Rect{500, 410, 530, 450}
	answer4[0] = false
	answer4[1] = true
	answer4[2] = true
	answer4[3] = false

	for i := 0; i < 4; i++ {
		got := FallsFromBlock(sample_test4[i])
		want := answer4[i]
		if got != want {
			t.Error("Falls From Block was incorrect, got: ", got, "want: ", want)
		}
	}
}
