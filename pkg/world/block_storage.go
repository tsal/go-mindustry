package world

import (
	"go-mindustry/pkg/game"
	"go-mindustry/pkg/types"
	"math"
)

type BlockStorage struct {
	game.UnlockableContent
	hasItems, hasLiquids, hasPower bool
	OutputsLiquid                  bool
	ConsumesPower                  bool
	OutputsPower                   bool

	itemCapacity   int
	LiquidCapacity float64

	stats, bars, consumes byte // placeholder fields
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L32
}

var _ types.Content = new(BlockStorage)

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

func (u *Unit) getTeam() int {
	// TODO: implement
	return 0
}

func (b BlockStorage) AcceptStack(item *Item, amount int, tile *Tile, source *Unit) (stack int) {
	stack = 0
	if b.AcceptItem(*item, *tile, *tile) && b.hasItems && (source == nil || source.getTeam() == tile.getTeam()) {
		// TODO: Implement tile item subtraction
		// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L51
		stack = int(math.Min(float64(b.GetMaximumAccepted(*tile, *item)), float64(amount)))
	}
	return
}
func (b *BlockStorage) RemoveStack(tile *Tile, item *Item, amount int) (stack int) {
	stack = 0
	if tile.entity == nil {
		return
	}
	// TODO: implement entity methods
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L63
	stack = int(math.Min(float64(amount), 0))
	return
}

func (b *BlockStorage) HandleStack(item *Item, amount int, tile *Tile, source *Unit) {
	// TODO: implement entity methods
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L72
	return
}

func (b BlockStorage) OutputsItems() bool {
	return b.hasItems
}

func (b BlockStorage) GetStackOffset(item *Item, tile *Tile, vector []int) {
	return
}

func (b *BlockStorage) OnProximityUpdate(tile *Tile) {
	if tile.entity != nil {
		// TODO: implement entity methods
		// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L86
	}
}

func (b *BlockStorage) HandleItem(item *Item, tile *Tile, source *Tile) {
	// TODO: implement entity methods
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L90
}

func (b *BlockStorage) HandleLiquid(tile *Tile, tile2 *Tile, liquid Liquid, amount float64) {
	// TODO: implement entity methods
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/world/BlockStorage.java#L102
}

func (b *BlockStorage) TryDumpLiquid(*Tile, Liquid) {
	panic("implement me!")
}

func (b BlockStorage) GetMaximumAccepted(Tile, Item) int {
	return b.itemCapacity
}

func (BlockStorage) AcceptItem(Item, Tile, Tile) bool {
	// TODO: implement Consumer type
	return true
}

type Liquid struct{}

func (BlockStorage) AcceptLiquid(Item, Tile, Liquid, amount float64) bool {
	// TODO: implement Consumer type
	return true
}
