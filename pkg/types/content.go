package types

import "fmt"

type ContentType int

const (
	Item ContentType = iota
	Block
	Mech
	Bullet
	Liquid
	Status
	Unit
	Weather
	Effect
	Zone
	LoadOut
	TypeId
)

type Content interface {
	ID() int32
	Content()
	ContentType() ContentType
	Init()
	Load()
	fmt.Stringer
}

type MappableContent interface {
	Content
	Name() string
	SetName(string)
}
