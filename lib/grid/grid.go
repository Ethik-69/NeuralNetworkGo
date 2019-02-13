package grid

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Grid struct {
	Height           int
	Width            int
	ChanceFood       float32
	ChangeRandomFood float32
	Grid             [][]interface{}
}

func (g *Grid) Init(height int, width int, chanceFood float32, ChangeRandomFood float32) {
	g.Height = height
	g.Width = width
	g.ChanceFood = chanceFood
	g.ChangeRandomFood = ChangeRandomFood
	g.Grid = make([][]interface{}, height)
}

func (g *Grid) RandomGen() {
	// Randomize the grid
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < g.Height; i++ {
		var buffer []interface{}
		for j := 0; j < g.Width; j++ {
			if rand.Float32() < g.ChanceFood {
				buffer = append(buffer, 1)
			} else {
				buffer = append(buffer, 0)
			}
		}
		g.Grid[i] = buffer
	}
}

func (g *Grid) AddRandomFood() {
	// Randomly add food on the grid
	for i := 0; i < 10; i++ {
		if rand.Float32() < g.ChangeRandomFood {
			x := rand.Intn(g.Width - 1)
			y := rand.Intn(g.Height - 1)
			if g.Grid[x][y] == 0 {
				g.Grid[x][y] = 1
			}
		}
	}
}

func (g *Grid) Update() {
	g.AddRandomFood()
}

func (g *Grid) cleanScreen() {
	// Launch the cmd clear
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Grid) PrettyPrint() {
	g.cleanScreen()
	var buffer bytes.Buffer
	for i := 0; i < g.Width; i++ {
		for j := 0; j < g.Height; j++ {
			switch g.Grid[i][j] {
			case 1:
				buffer.WriteString("X")
			case 0:
				buffer.WriteString(" ")
			default:
				buffer.WriteString("\u2B1C")
			}
		}
		buffer.WriteString("\n")
	}
	fmt.Printf("\033[2K\r%s", buffer.String())
}
