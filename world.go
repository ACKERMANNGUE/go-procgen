package main

import (
	"fmt"
	"strings"
)

type World struct {
	Width  int
	Height int
	Tiles []Tile
}

func NewWorld(width, height int) *World {
	tiles := make([]Tile, width*height)

	return &World{
		Width:  width,
		Height: height,
		Tiles:  tiles,
	}
}

func (w *World) index(x, y int) int {
	return y*w.Width + x
}

func (w *World) IsInside(x, y int) bool {
	return x >= 0 && x < w.Width && y >= 0 && y < w.Height
}

func (w *World) Get(x, y int) Tile {
	if !w.IsInside(x, y) {
		return TileWall
	}

	return w.Tiles[w.index(x, y)]
}

func (w *World) Set(x, y int, tile Tile) {
	if !w.IsInside(x, y) {
		return
	}

	w.Tiles[w.index(x, y)] = tile
}

func (w *World) String() string {
	var sb strings.Builder
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			tile := w.Get(x, y)
			sb.WriteRune(tile.Rune())
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (w *World) Print() {
	fmt.Print(w.String())
}