package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"math"
	"sort"
	//"log"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
	"github.com/beefsack/go-astar"
)

const STARTING_HITPOINTS int = 200
const ATTACK int = 3

func main() {
	lines := utils.GetLines("../myInput.txt")
	//lines := utils.GetLines("../test1.txt")

	cave := World{}
	units := make([]*Unit, 0)

	//Load in cave and units
	for y, line := range lines {
		for x, cell := range line {
			if cell == '#' || cell == '.' {
				cave.SetTile(&Tile{
					kind:cell,
					x:x,
					y:y,
				})
			} else {
				units = append(units, &Unit{
					race:      cell,
					x:         x,
					y:         y,
					hitpoints: STARTING_HITPOINTS,
				})
				cave.SetTile(&Tile{
					kind:'.',
					x:x,
					y:y,
					unit:units[len(units)-1],
				})
			}
		}
	}
	mapPrinter(cave)

	round := 0
	targetsAvailable := true
	roundCap := 1000
	for targetsAvailable && round < roundCap {
		targetsAvailable = false
		round++

		//for _,unit := range units {
		//	unitPrinter(unit)
		//}
		//fmt.Println()
		sort.Sort(Units(units))
		//for _,unit := range units {
		//	unitPrinter(unit)
		//}
		
		for i := 0; i < len(units); i++ {
			currentUnit := units[i]
			if currentUnit.hitpoints > 0 { //check selected unit is not dead
				nearestTarget, targetDistance,targets := findNearest(currentUnit, units)
				if nearestTarget != nil {
					targetsAvailable = true
					if targetDistance == 1 { //attack
						attack(nearestTarget,cave)
					} else { //move
						move(currentUnit,cave,targets)
					}
				}
			}
		}
		time.Sleep(time.Second/10)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		fmt.Println("Round:",round)
		mapPrinter(cave)
		//for _,unit := range units {
		//	unitPrinter(unit)
		//}
		fmt.Println()
		
	}
	sum:= 0
	for _, unit := range units {
		if unit.hitpoints > 0{
		sum = sum + unit.hitpoints
		}
	}
	fmt.Println("Sum:", sum,"Round:", round,"Outcome:",sum*(round))
	
}

func move(currentUnit *Unit, cave World, targets []*Unit){
	open := make([]*Tile,0)
	for i := range targets {
		t := targets[i]

		//identify open squares in range of target
		for _, offset := range [][]int{
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			if n := cave.Tile(t.x+offset[0],t.y+offset[1]); 
				n != nil &&
				n.kind != '#' &&
				n.unit == nil {
				open = append(open, n)
			}
		}

	}

	//which squares in fewest steps
	var shortPath []astar.Pather
	var compare float64 = 1000
	for i := range open {
		path, distance, found := astar.Path(cave.Tile(currentUnit.x,currentUnit.y),open[i])
		if found && distance < compare {
			shortPath = path
			compare = distance
		}
	}
	if compare < 1000 {
	//move one space on shortest path
	toT := shortPath[len(shortPath)-2].(*Tile)

	cave.Tile(currentUnit.x,currentUnit.y).unit = nil
	currentUnit.x = toT.x
	currentUnit.y = toT.y
	cave.Tile(currentUnit.x,currentUnit.y).unit = currentUnit
	}
}

func attack(target *Unit,cave World){
	target.hitpoints = target.hitpoints - 3
	if target.hitpoints <= 0 {
		cave.Tile(target.x,target.y).unit = nil
	}
}

func findNearest(currentUnit *Unit, units []*Unit) (*Unit, int, []*Unit) {
	//find targets
	var nearestTarget *Unit
	targetDistance := int(^uint(0) >> 1)
	targets := make([]*Unit,0)
	for j := range units {
		possibleTarget := units[j]
		if currentUnit.race != possibleTarget.race && //check if different race
			possibleTarget.hitpoints > 0 { //check if not dead
			targets = append(targets,possibleTarget)
			possibleDistance := distanceBetween(currentUnit, possibleTarget)
			if possibleDistance < targetDistance ||
			  (possibleDistance == targetDistance && nearestTarget.hitpoints > possibleTarget.hitpoints) {
				nearestTarget = possibleTarget
				targetDistance = possibleDistance
			}
		}
	}
	return nearestTarget, targetDistance, targets
}

func distanceBetween(unit *Unit, unit1 *Unit) int {
	return int(math.Abs(float64(unit.x-unit1.x)) + math.Abs(float64(unit.y-unit1.y)))
}

func unitPrinter(unit *Unit) {
	if unit.hitpoints > 0 {
	fmt.Print("{")
	fmt.Print(string(unit.race), " ")
	fmt.Print("x:", unit.x, " y:", unit.y, " ")
	fmt.Print(unit.hitpoints, "}")
	}
}

func mapPrinter(cave World) {

	h:= len(cave)
	w:= len(cave[0])

	for y:=0; y < h; y++  {
		units := make([]*Unit,0)
		for x:=0; x<w; x++ {
			tile := cave.Tile(x,y)
			if tile.unit == nil || tile.unit.hitpoints <= 0 {
				fmt.Print(string(tile.kind))
			} else {
				fmt.Print(string(tile.unit.race))
				units = append(units,tile.unit)
			}
		}
		fmt.Print("  ")
		for _,unit := range units {
			unitPrinter(unit)
		}
		fmt.Println()
	}
}


type World map[int]map[int]*Tile

type Tile struct{
	kind rune
	cave World
	unit *Unit
	x int
	y int
}

type Unit struct {
	race      rune
	x         int
	y         int
	hitpoints int
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
			n.kind != '#' && n.unit == nil{
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



//Unit slice type and sorting functions for sort.Sort interface
type Units []*Unit

func (s Units) Len() int {
	return len(s)
}

func (s Units) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Units) Less(i, j int) bool {
	return s[i].y*10 + s[i].x < s[j].y*10 + s[j].x
}