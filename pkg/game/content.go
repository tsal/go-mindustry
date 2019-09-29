package game

import "go-mindustry/pkg/types"

type MappableContent struct {
	Name string
}

func (m MappableContent) ID() int32 {
	panic("implement me")
}

func (m MappableContent) Content() {
	panic("implement me")
}

func (m MappableContent) ContentType() types.ContentType {
	panic("implement me")
}

func (m MappableContent) Init() {
	panic("implement me")
}

func (m MappableContent) Load() {
	panic("implement me")
}

func (m MappableContent) String() string {
	return m.Name
}

func NewMappableContent(name string) (C *MappableContent) {
	C = new(MappableContent)
	C.Name = name
	return
}

var _ types.Content = NewMappableContent("")

type UnlockableContent struct {
	*MappableContent
	localizedName string
	Description   string
}

func NewUnlockableContent(name string) (C *UnlockableContent) {
	C = new(UnlockableContent)
	C.MappableContent = NewMappableContent(name)
	// TODO: Implement localization
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/game/UnlockableContent.java#L18
	return
}

func (u UnlockableContent) LocalizedName() string {
	return u.localizedName
}

func (u UnlockableContent) OnUnlock()            {}
func (u UnlockableContent) Hidden() bool         { return false }
func (u UnlockableContent) AlwaysUnlocked() bool { return false }
func (u *UnlockableContent) Unlocked() bool {
	// TODO: Implement Vars
	// https://github.com/tsal/Mindustry/blob/master/core/src/io/anuke/mindustry/game/UnlockableContent.java#L44
	return true
}
func (u UnlockableContent) Locked() bool { return !u.Unlocked() }

var _ types.Content = NewUnlockableContent("")
