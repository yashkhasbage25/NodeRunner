package dijkstra

import (
	"testing"

	"github.com/IITH-SBJoshi/concurrency-3/src/dtypes"
	"github.com/IITH-SBJoshi/concurrency-3/src/platform"
)

func TestNextposition(t *testing.T) {
	var got [5]dtypes.Position
	var want [5]dtypes.Position
	got[0] = nextposition(dtypes.Position{42, 89}, dtypes.Position{42, 114}, 2)
	got[1] = nextposition(dtypes.Position{568, 456}, dtypes.Position{568, 200}, 2)
	got[2] = nextposition(dtypes.Position{312, 124}, dtypes.Position{78, 124}, 2)
	got[3] = nextposition(dtypes.Position{245, 67}, dtypes.Position{300, 67}, 2)
	got[4] = nextposition(dtypes.Position{67, 89}, dtypes.Position{67, 89}, 2)
	want[0] = dtypes.Position{42, 91}
	want[1] = dtypes.Position{568, 454}
	want[2] = dtypes.Position{310, 124}
	want[3] = dtypes.Position{247, 67}
	want[4] = dtypes.Position{67, 89}
	for i := 0; i < 5; i++ {

		if got[i] != want[i] {
			t.Error("nextposition was incorrect got :", got[i], "want ", want[i])
		}
	}
}

func TestOnPlatform(t *testing.T) {
	platform.Initialize()
	var got [5]int
	var want [5]int
	got[0] = OnPlatform(GetBoundary(dtypes.Position{140, 180}))
	got[1] = OnPlatform(GetBoundary(dtypes.Position{900, 420}))
	got[2] = OnPlatform(GetBoundary(dtypes.Position{20, 78}))
	got[3] = OnPlatform(GetBoundary(dtypes.Position{245, 67}))
	got[4] = OnPlatform(GetBoundary(dtypes.Position{67, 89}))
	want[0] = 1
	want[1] = 1
	want[2] = 0
	want[3] = 0
	want[4] = 0
	for i := 0; i < 5; i++ {

		if got[i] != want[i] {
			t.Error("nextposition was incorrect got :", got[i], "want ", want[i])
		}
	}
}
func TestMinimum(t *testing.T) {
	var got [5]int
	var want [5]int
	distance := []int{23, 34, 18, 56, 72, 10, 112}
	got[0] = minimum(distance, 1, 2)
	got[1] = minimum(distance, 5, 6)
	got[2] = minimum(distance, 0, 4)
	got[3] = minimum(distance, 4, 2)
	got[4] = minimum(distance, 5, 3)
	want[0] = 2
	want[1] = 5
	want[2] = 0
	want[3] = 2
	want[4] = 5
	for i := 0; i < 5; i++ {

		if got[i] != want[i] {
			t.Error("nextposition was incorrect got :", got[i], "want ", want[i])
		}
	}
}
