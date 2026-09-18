package main

type Tile int

const (
	TileGround Tile = iota // iota assigns successive integer values starting from 0
	TileWall
	TileWater
)

func (t Tile) Rune() rune {
	switch t {
	case TileGround:
		return '.'
	case TileWall:
		return '#'
	case TileWater:
		return '~'
	default:
		return '?'
	}
}
