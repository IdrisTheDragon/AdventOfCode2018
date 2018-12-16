package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
	//"log"

	"github.com/IdrisTheDragon/AdventOfCode2018/utils"
	"github.com/beefsack/go-astar"
)

const STARTING_HITPOINTS int = 200
const ATTACK int = 3

func main() {
	lines := utils.GetLines("../myInput.txt")
	//lines := utils.GetLines("../testM.txt")

	cave := World{}
	units := make([]Unit, 0)

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
				cave.SetTile(&Tile{
					kind:'.',
					x:x,
					y:y,
				})
				units = append(units, Unit{
					race:      cell,
					x:         x,
					y:         y,
					hitpoints: STARTING_HITPOINTS,
				})
			}
		}
	}

	mapPrinter(cave, units)

	round := 0
	targetsAvailable := true
	roundCap := 1000
	for targetsAvailable && round < roundCap {
		targetsAvailable = false
		round++
		for i := 0; i < len(units); i++ {
			currentUnit := &units[i]
			if currentUnit.hitpoints > 0 { //check selected unit is not dead
				nearestTarget, targetDistance := findNearest(currentUnit, &units)
				if nearestTarget != nil {
					targetsAvailable = true
					if targetDistance == 1 { //attack
						attack(nearestTarget)
					} else { //move
						move(currentUnit,nearestTarget,cave)
					}
				}
			}
		}
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		mapPrinter(cave, units)
		//for _,unit := range units {
			//unitPrinter(&unit)
		//}
		fmt.Println(round)
		time.Sleep(time.Second/5)
	}
	sum:= 0
	for _, unit := range units {
		if unit.hitpoints > 0{
		sum = sum + unit.hitpoints
		}
	}
	fmt.Println(sum*(round-1))
	//mapPrinter(cave, units)
	
}

func move(currentUnit *Unit, target *Unit, cave World){
	//pathfinding..
	t1 := cave.Tile(currentUnit.x,currentUnit.y)
	t2 := cave.Tile(target.x,target.y)
	path, _, found := astar.Path(t1, t2)
	if !found {
		//log.Println("Could not find path")
	} else {
		toT := path[len(path)-2].(*Tile)
		//fmt.Println(toT.x,toT.y)
		currentUnit.x = toT.x
		currentUnit.y = toT.y
	}
}

func attack(target *Unit){
	target.hitpoints = target.hitpoints - 3
}

func findNearest(currentUnit *Unit, units *[]Unit) (*Unit, int) {
	//find targets
	var nearestTarget *Unit
	targetDistance := int(^uint(0) >> 1)
	for j := range *units {
		possibleTarget := &(*units)[j]
		if currentUnit.race != possibleTarget.race && //check if different race
			possibleTarget.hitpoints > 0 { //check if not dead
			possibleDistance := distanceBetween(currentUnit, possibleTarget)
			if possibleDistance < targetDistance {
				nearestTarget = possibleTarget
				targetDistance = possibleDistance

			}
		}
	}
	return nearestTarget, targetDistance
}

func distanceBetween(unit *Unit, unit1 *Unit) int {
	return int(math.Abs(float64(unit.x-unit1.x)) + math.Abs(float64(unit.y-unit1.y)))
}

func unitPrinter(unit *Unit) {
	fmt.Print("{")
	fmt.Print(string(unit.race), " ")
	fmt.Print("x:", unit.x, " y:", unit.y, " ")
	fmt.Print(unit.hitpoints, "}")
}

func mapPrinter(cave World, units []Unit) {

	for i, _ := range units {
		tile := cave.Tile(units[i].x,units[i].y)
		tile.unit = &units[i]
	}

	h:= len(cave)
	w:= len(cave[0])

	for y:=0; y < h; y++  {
		for x:=0; x<w; x++ {
			tile := cave.Tile(x,y)
			if tile.unit == nil {
				fmt.Print(string(tile.kind))
			} else {
				fmt.Print(string(tile.unit.race))
				tile.unit = nil
			}
		}
		fmt.Println()
	}
}

type Unit struct {
	race      rune
	x         int
	y         int
	hitpoints int
}
