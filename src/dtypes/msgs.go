package dtypes

// Position represents the coordinates of an object in the form (left, top)
type Position struct {
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
}

// Event represents the events on the client side
type Event struct {
	EventType string `json:"etype"`

	Object string `json:"object"`

	P1Pos Position `json:"p1_pos"`
	P2Pos Position `json:"p2_pos"`

	B1Pos Position `json:"b1_pos"`
	B2Pos Position `json:"b2_pos"`
	B3Pos Position `json:"b3_pos"`

	G1Pos Position `json:"g1_pos"`
	G2Pos Position `json:"g2_pos"`
	G3Pos Position `json:"g3_pos"`
	G4Pos Position `json:"g4_pos"`
}
