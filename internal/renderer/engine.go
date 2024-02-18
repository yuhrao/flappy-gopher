package renderer

import "github.com/charmbracelet/lipgloss"
// import "math/rand/v2"

const (
	gopherPosX        = 10
	maxObstacleWidth  = 5
	maxObstaclesCount = 5
)

var (
	wallChar = '█'
	bgChar   = '▒'
)

type Engine struct {
	WindowSize [2]int
	BirdHeight int
	Obstacles  []*Obstacle
}

// New Engine
func NewEngine(wSize [2]int) *Engine {
	birdHeight := int(wSize[1] / 2)
	obstacles := make([]*Obstacle, 0)
	eng := &Engine{wSize, birdHeight, obstacles}

	eng.initialize()

	return eng
}

func (e *Engine) initialize() {
	for i := 0; i < maxObstaclesCount; i++ {
		e.AddObstacle()
	}
}

func (e *Engine) AddObstacle() {
	if len(e.Obstacles) >= maxObstaclesCount {
		return
	}
	width := randomIntBetween(1, maxObstacleWidth)
	gap := randomIntBetween(2, 4)
	quarterHeight, halfHeight := int(e.WindowSize[1]/4), int(e.WindowSize[1]/2)
	height := randomIntBetween(quarterHeight, halfHeight+quarterHeight)
	px := e.WindowSize[0]

	if len(e.Obstacles) > 0 {
		lastObstacle := e.Obstacles[len(e.Obstacles)-1]
		distanceBetweenObstacles := randomIntBetween(1, 10)
		px = lastObstacle.px + lastObstacle.width + distanceBetweenObstacles
	}

	newObstacle := NewObstacle(px, gap, height, width)
	e.Obstacles = append(e.Obstacles, newObstacle)
}

func (e *Engine) Move() {
  for i, o := range e.Obstacles {
    o.Move(e.WindowSize[0])
    e.Obstacles[i] = o
  }
}


func (e *Engine) isObstacle(x, y int) bool {
	for _, o := range e.Obstacles {
		if o.IntersectingY(y) && o.IntersectingX(x) {
			return true
		}
	}
	return false
}

func (e *Engine) getRune(px, py int) rune {

	if e.isObstacle(px, py) {
		return wallChar
	}

	if py == e.BirdHeight && px == gopherPosX {
		return 'G'
	}

	return bgChar

}

func (e *Engine) createCanvas() [][]rune {
	canvas := make([][]rune, e.WindowSize[1])
	for py, x := range canvas {
		x = make([]rune, e.WindowSize[0])
		for px := range x {
			x[px] = e.getRune(px, py)
		}
		canvas[py] = x
	}
	return canvas
}

func (e *Engine) Render() string {
	s := ""
	for _, row := range e.createCanvas() {
		for _, cell := range row {
			var style lipgloss.Style
			if cell == wallChar {
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF9580")).Background(lipgloss.Color("#FF9580"))
			} else if bgChar == cell {
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("#414D58"))
			} else {
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(lipgloss.Color("#414D58"))
			}
			s += style.Render(string(cell))
		}
		s += "\n"
	}
	return s
}
