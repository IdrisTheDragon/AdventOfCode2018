package main

type World map[int]map[int]*Tile

type Tile struct{
	kind rune
	cave World
	unit *Unit
	x int
	y int
}

func (w World) SetTile(t *Tile) {
	if w[t.y] == nil {
		w[t.y] = map[int]*Tile{}
	}
	w[t.y][t.x] = t
	t.cave = w
}

func (w World) Tile(x, y int) *Tile {
	if w[y] == nil {
		return nil
	}
	return w[y][x]
}