package main

import "github.com/beefsack/go-astar"

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbours := []astar.Pather{}
	cave := t.cave
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := cave.Tile(t.x+offset[0],t.y+offset[1]); n != nil &&
			n.kind != '#' {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}