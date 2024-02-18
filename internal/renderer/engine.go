package renderer

import (
	"strconv"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	log "github.com/yuhrao/flappy-gopher/internal"
)

const (
	gopherPosX        = 10
	maxObstacleWidth  = 5
	maxObstaclesCount = 10
)

var (
	wallChar = '█'
	bgChar   = '▒'
)

type Engine struct {
	WindowSize [2]int
	BirdHeight int
	Obstacles  []*Obstacle
	mux        sync.Mutex
}

// New Engine
func NewEngine(wSize [2]int) *Engine {
	birdHeight := int(wSize[1] / 2)
	obstacles := make([]*Obstacle, 0)
	mux := sync.Mutex{}
	eng := &Engine{wSize, birdHeight, obstacles, mux}

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
	gap := randomIntBetween(4, 10)
	quarterHeight, halfHeight := int(e.WindowSize[1]/4), int(e.WindowSize[1]/2)
	height := randomIntBetween(quarterHeight, halfHeight+quarterHeight)
	var px int
	obstacleCount := len(e.Obstacles)

	if obstacleCount > 0 {
		lastObstacle := e.Obstacles[len(e.Obstacles)-1]
		distanceBetweenObstacles := randomIntBetween(8, 30)
		px = lastObstacle.px + lastObstacle.width + distanceBetweenObstacles
	} else {
		px = e.WindowSize[0]
	}

	newObstacle := NewObstacle(px, gap, height, width)
	e.Obstacles = append(e.Obstacles, newObstacle)
}

func (e *Engine) Move() {
	e.mux.Lock()

	logFields := logrus.Fields{}

	for i, o := range e.Obstacles {
		logFields[strconv.Itoa(i)] = o.px

		o.Move(e.WindowSize[0])
		e.Obstacles[i] = o
	}

	log.Logger.WithFields(logFields).Info("Moving obstacle")

	firstObstacle := e.Obstacles[0]
	if firstObstacle.px < -firstObstacle.width {
		e.Obstacles = e.Obstacles[1:]
		e.AddObstacle()
	}

	e.mux.Unlock()
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
	e.mux.Lock()
	defer e.mux.Unlock()

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
