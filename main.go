package main

func main() {
	world := NewWorld(40, 15)

	addBorders(world)
	addRectangle(
		world,
		10,
		3,
		8,
		4,
		TileWall,
	)
	addRectangle(
		world,
		23,
		8,
		6,
		3,
		TileWater,
	)

	world.Print()
}

func addBorders(world *World) {
	for x := 0; x < world.Width; x++ {
		world.Set(x, 0, TileWall)
		world.Set(x, world.Height-1, TileWall)
	}

	for y := 0; y < world.Height; y++ {
		world.Set(0, y, TileWall)
		world.Set(world.Width-1, y, TileWall)
	}
}

func addRectangle(
	world *World,
	startX int,
	startY int,
	width int,
	height int,
	tile Tile,
) {
	for y := startY; y < startY+height; y++ {
		for x := startX; x < startX+width; x++ {
			world.Set(x, y, tile)
		}
	}
}
