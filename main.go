package main

import (
	"game_of_life_with_NN/lib/cell"
	"game_of_life_with_NN/lib/grid"
	"sort"
	"time"
)

const (
	height           = 50
	width            = 50
	chanceFood       = 0.1
	chanceFoodRandom = 0.1
	maxCells         = 20
	numInput         = 2
	numHidden        = 4
	numOutput        = 4
)

func addCells(grid grid.Grid) []cell.Cell {
	// Create all cells
	var cells []cell.Cell
	for {
		if len(cells) == maxCells {
			return cells
		}
		newCell := new(cell.Cell)
		newCell.Init(grid, numInput, numHidden, numOutput)
		cells = append(cells, *newCell)
	}
	return cells
}

func updateCells(cells []cell.Cell) {
	for i := 0; i < len(cells); i++ {
		cells[i].Update()
	}

}

func genAlgo(cells []cell.Cell) {
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].Fitness < cells[j].Fitness
	})

	cells = cells[:maxCells/3]

	for i := 0; i < len(cells); i++ {
		cells[i].Fitness = 0
	}

}

func main() {
	grid := new(grid.Grid)
	grid.Init(height, width, chanceFood, chanceFoodRandom)
	grid.RandomGen()
	cells := addCells(*grid)
	for {
		for i := 0; i < 500; i++ {
			updateCells(cells)
			grid.Update()
			grid.PrettyPrint()
			time.Sleep(20 * time.Millisecond)
		}
		//genAlgo(cells)
	}
}
