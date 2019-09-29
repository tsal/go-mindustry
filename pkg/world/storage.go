package world

import "go-mindustry/pkg/game"

type BlockStorage struct {
	*game.UnlockableContent
	hasItems, hasLiquids, hasPower bool
	OutputsLiquid                  bool
	ConsumesPower                  bool
	OutputsPower                   bool

	itemCapacity   int
	LiquidCapacity float64

	stats, bars, consumes byte // placeholder fields
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L32
}

func (b BlockStorage) ShouldConsume(T Tile) bool {
	return true
}

func (b BlockStorage) GetPowerProduction(T Tile) float64 {
	return 0
}

// Item implements a placeholder
// TODO: Implement in the appropriate place
type Item struct{}

// Unit implements a placeholder
// TODO: Implement in the appropriate place
type Unit struct{}

func (b BlockStorage) AcceptStack() (stack int) {
	stack = 0

	return
}

func (b BlockStorage) getMaximumAccepted(tile Tile, item Item) int {
	return b.itemCapacity
}

func (BlockStorage) acceptItem(item Item, tile Tile, source Tile) bool {
	// TODO: implement Consumer type
	return true
}
