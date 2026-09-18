package main

type Tile int

const (
	TileGround Tile = iota // iota assigns successive integer values starting from 0
	TileWall
	TileWater
	TileMountain
	TileForest
	TileSand
	TileLava
	TileIce
	TileSwamp
	TileBridge
	TileDoor
	TileGrass
	TileSnow
	TileRoad
	TileCliff
	TileCave
)

func (t Tile) Rune() rune {
	switch t {
	case TileGround:
		return '.'
	case TileWall:
		return '#'
	case TileWater:
		return '~'
	case TileMountain:
		return '^'
	case TileForest:
		return '*'
	case TileSand:
		return ':'
	case TileLava:
		return '!'
	case TileIce:
		return 'o'
	case TileSwamp:
		return '%'
	case TileBridge:
		return '='
	case TileDoor:
		return '+'
	case TileGrass:
		return '"'
	case TileSnow:
		return '\''
	case TileRoad:
		return '-'
	case TileCliff:
		return '|'
	case TileCave:
		return '@'
	default:
		return '?'
	}
}
