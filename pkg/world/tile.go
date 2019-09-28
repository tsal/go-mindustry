package world

type Position struct{}
type TargetTrait struct{}

type Tile struct {
	*Position
	*TargetTrait
}
