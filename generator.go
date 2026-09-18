package main

import "math/rand"

type Generator struct {
	random *rand.Rand
}

func NewGenerator(seed int64) *Generator {
	return &Generator{
		random: rand.New(rand.NewSource(seed)),
	}
}

func (g *Generator) Generate(world *World) {
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			tile := g.randomTile()
			world.Set(x, y, tile)
		}
	}
}

func (g *Generator) generateNoise(world *World) {
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			tile := g.randomTile()
			world.Set(x, y, tile)
		}
	}
}

func (g *Generator) randomTile() Tile {
	value := g.random.Float64()
	switch {
	case value < 0.30:
		return TileGrass
	case value < 0.55:
		return TileWater
	case value < 0.75:
		return TileForest
	case value < 0.90:
		return TileMountain
	default:
		return TileGround
	}
}

func (g *Generator) smooth(world *World, iterations int) {
	for range iterations {
		g.smoothOnce(world)
	}
}

func (g *Generator) smoothOnce(world *World) {
	newTile := make([]Tile, world.Width*world.Height)
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			newTile[world.index(x, y)] = g.mostCommonNeighbour(world, x, y)
		}
	}

	world.Tiles = newTile
}

func (g *Generator) mostCommonNeighbour(world *World, x, y int) Tile {
	counts := make(map[Tile]int)
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -1; offsetX <= 1; offsetX++ {
			neighbourX := x + offsetX
			neighbourY := y + offsetY

			if !world.IsInside(neighbourX, neighbourY) {
				continue
			}

			tile := world.Get(neighbourX, neighbourY)
			counts[tile]++
		}
	}

	currentTile := world.Get(x, y)
	bestTile := currentTile
	bestCount := counts[currentTile]

	for tile, count := range counts {
		if count > bestCount {
			bestTile = tile
			bestCount = count
		}
	}

	return bestTile
}