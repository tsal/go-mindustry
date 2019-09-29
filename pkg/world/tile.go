package world

// Position is a placeholder type
// TODO: Implement Position in appropriate place
type Position struct{}

// TargetTrait is a placeholder type
// TODO: Implement TargetTrait in appropriate place
type TargetTrait struct{}

// TileEntity is a placeholder type
// TODO: Implement TileEntity in appropriate place
type TileEntity struct{}

// Tile is a WIP type
type Tile struct {
	*Position
	*TargetTrait
	Cost                  byte
	Entity                TileEntity
	X, Y                  int32
	block, floor, overlay byte // placeholder types
	rotation, team        byte
}

func New(x, y int) (T Tile) {
	T = Tile{
		Position:    new(Position),
		TargetTrait: new(TargetTrait),
		Cost:        1,
		X:           int32(x),
		Y:           int32(y),
	}

	return
}
