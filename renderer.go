package main

import (
	"fmt"
	"strings"
)

const (
	ColorReset = "\033[0m"
	ColorGray   = "\033[90m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorMagenta = "\033[35m"
	ColorBlack  = "\033[30m"
	ColorOrange = "\033[38;5;208m"
	ColorPurple = "\033[38;5;93m"
	ColorBrown  = "\033[38;5;94m"
	ColorPink   = "\033[38;5;213m"
	ColorTeal   = "\033[38;5;6m"
	ColorLime   = "\033[38;5;10m"
	ColorIndigo = "\033[38;5;54m"
	ColorViolet = "\033[38;5;129m"
)

func tileColor(tile Tile) string {
	switch tile {
	case TileGround:
		return ColorGreen
	case TileWall:
		return ColorGray
	case TileWater:
		return ColorBlue
	case TileMountain:
		return ColorGray
	case TileForest:
		return ColorGreen
	case TileSand:
		return ColorYellow
	case TileLava:
		return ColorRed
	case TileIce:
		return ColorCyan
	case TileSwamp:
		return ColorGreen
	case TileBridge:
		return ColorBrown
	case TileDoor:
		return ColorIndigo
	case TileGrass:
		return ColorGreen
	case TileSnow:
		return ColorWhite
	case TileRoad:
		return ColorGray
	case TileCliff:
		return ColorBrown
	case TileCave:
		return ColorBlack
	default:
		return ColorReset
	}
}

func RenderWorld(world *World) string {
	var sb strings.Builder

	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			tile := world.Get(x, y)

			sb.WriteString(tileColor(tile))
			sb.WriteRune(tile.Rune())
			sb.WriteString(ColorReset)
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}

func PrintWorld(world *World) {
	fmt.Print(RenderWorld(world))
}