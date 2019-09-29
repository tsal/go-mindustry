package game

import "C"
import "go-mindustry/pkg/types"

type gameUnlockableContent struct{}

func (g *gameUnlockableContent) ID() int32 {
	panic("implement me")
}

func (g *gameUnlockableContent) Content() {
	panic("implement me")
}

func (g *gameUnlockableContent) ContentType() types.ContentType {
	panic("implement me")
}

func (g *gameUnlockableContent) Init() {
	panic("implement me")
}

func (g *gameUnlockableContent) Load() {
	panic("implement me")
}

func (g *gameUnlockableContent) String() string {
	panic("implement me")
}

func (g *gameUnlockableContent) Name() string {
	panic("implement me")
}

func (g *gameUnlockableContent) SetName(string) {
	panic("implement me")
}

func (g *gameUnlockableContent) LocalizedName() string {
	panic("implement me")
}

func (g *gameUnlockableContent) OnUnlock() {
	panic("implement me")
}

func (g *gameUnlockableContent) Hidden() bool {
	panic("implement me")
}

func (g *gameUnlockableContent) AlwaysUnlocked() bool {
	panic("implement me")
}

func (g *gameUnlockableContent) Unlocked() bool {
	panic("implement me")
}

func (g *gameUnlockableContent) Locked() bool {
	panic("implement me")
}

func NewUnlockableContent(name string) (C UnlockableContent) {
	C = new(gameUnlockableContent)
	// TODO: Implement localization
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/game/UnlockableContent.java#L18
	return
}

type UnlockableContent interface {
	types.MappableContent
	LocalizedName() string
	OnUnlock()
	Hidden() bool
	AlwaysUnlocked() bool
	Unlocked() bool
	Locked() bool
}
