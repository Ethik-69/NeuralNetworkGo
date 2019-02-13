package cell

import (
	"game_of_life_with_NN/lib/grid"
	"game_of_life_with_NN/lib/network"
	"math/rand"
	"time"
)

type Cell struct {
	X, Y        int
	Alive       bool
	Fitness     int
	found       [][]float64
	Sensor      [2]interface{}
	SensorRange int
	SensorLimit int
	Brain       network.Network
	NumUpdate   int
	ErrorLevel  float64
	NInput      int
	NHidden     int
	NOutput     int
	grid        grid.Grid
}

func (c *Cell) RandomPos() int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	for {
		c.X = rand.Intn(len(c.grid.Grid) - 1)
		c.Y = rand.Intn(len(c.grid.Grid) - 1)
		if c.grid.Grid[c.X][c.Y] == 0 {
			c.grid.Grid[c.X][c.Y] = c
			return 1
		}
	}
}

func (c *Cell) makeRange(min, max int) []int {
	var output []int
	for i := min; i < max; i++ {
		if i < 0 {
			output = append(output, i+c.grid.Width)
		} else if i > c.grid.Width-1 {
			output = append(output, i-c.grid.Width)
		} else {
			output = append(output, i)
		}
	}
	return output
}

func (c *Cell) detection() {
	c.SensorRange = 1
	c.Sensor = [2]interface{}{0, 0}
	c.found = [][]float64{}
	sensorOn := true

	for sensorOn {
		xRange := c.makeRange(c.X-c.SensorRange, c.X+c.SensorRange+1)
		yRange := c.makeRange(c.Y-c.SensorRange, c.Y+c.SensorRange+1)
		for _, x := range xRange {
			for _, y := range yRange {
				if c.grid.Grid[x][y] == 1 {
					c.found = append(c.found, []float64{float64(x), float64(y)})
				}
			}
		}

		if c.SensorRange == c.SensorLimit {
			sensorOn = false
			number := 0.0
			sum := []float64{0.0, 0.0}
			for i := range c.found {
				sum[0] += c.found[i][0]
				sum[1] += c.found[i][1]
				number++
			}

			sum[0] = sum[0] / number
			sum[1] = sum[1] / number
			sensX := float32(int(sum[0])-c.X) / 10
			sensY := float32(int(sum[1])-c.Y) / 10
			c.Sensor = [2]interface{}{sensX, sensY}
		}
		c.SensorRange++
	}
}

func (c *Cell) eat() {
	if c.grid.Grid[c.X][c.Y] == 1 {
		c.grid.Grid[c.X][c.Y] = 0
		c.Fitness++
	}
}

func (c *Cell) move() {
	if c.Brain.ArrayOutput[0] > 0.5 {
		c.Y--
		if c.Y < 0 {
			c.Y = c.grid.Width - 1
		}
	}
	if c.Brain.ArrayOutput[1] > 0.5 {
		c.X++
		if c.X > c.grid.Width-1 {
			c.X = 0
		}
	}
	if c.Brain.ArrayOutput[2] > 0.5 {
		c.Y++
		if c.Y > c.grid.Width-1 {
			c.Y = 0
		}
	}
	if c.Brain.ArrayOutput[3] > 0.5 {
		c.X--
		if c.X < 0 {
			c.X = c.grid.Width - 1
		}
	}
}

func (c *Cell) Update() {
	c.NumUpdate++
	c.detection()
	c.Brain.Update(c.Sensor)

	targets := []float64{0.5, 0.5, 0.5, 0.5}

	if c.Sensor != [2]interface{}{0, 0} {
		if float64(c.Sensor[0].(float32)) < 0 {
			targets[1] = 0.4
			targets[3] = 0.6
		} else if float64(c.Sensor[0].(float32)) > 0 {
			targets[1] = 0.6
			targets[3] = 0.4
		}

		if float64(c.Sensor[1].(float32)) < 0 {
			targets[0] = 0.6
			targets[2] = 0.4
		} else if float64(c.Sensor[1].(float32)) > 0 {
			targets[0] = 0.4
			targets[2] = 0.6
		}
	}

	c.ErrorLevel = 0.0
	c.ErrorLevel = c.Brain.BackPropagation(targets)

	c.grid.Grid[c.X][c.Y] = 0
	c.move()
	c.eat()
	c.grid.Grid[c.X][c.Y] = c
}

func (c *Cell) Init(grid grid.Grid, numInput int, numHidden int, numOutput int) {
	c.grid = grid
	c.Alive = true
	c.Fitness = 0
	c.SensorRange = 1
	c.SensorLimit = 4
	c.NumUpdate = 0
	c.ErrorLevel = 0.0
	c.NInput = numInput
	c.NHidden = numHidden
	c.NOutput = numOutput
	c.RandomPos()
	c.Brain.Init(numInput, numHidden, numOutput)
}
