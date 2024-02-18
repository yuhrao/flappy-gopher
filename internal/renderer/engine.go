package renderer

const (
	birdPosX      = 10
	obstacleWidth = 3
	wallChar      = '█'
	bgChar        = '▒'
	emptyChar     = '~'
)


type RenderEngine struct {
	WindowSize [2]int
	BirdHeight int
	Obstacles  []Obstacle // pos, height
}

func (r RenderEngine) isObstacle(x, y int) bool {
	for _, o := range r.Obstacles {
		if o.IntersectingY(y) && o.IntersectingX(x) {
			return true
		}
	}
	return false
}

func (r RenderEngine) getRune(px, py int) rune {

	if r.isObstacle(px, py) {
		return wallChar
	}

	if py == r.BirdHeight && px == birdPosX {
		return 'B'
	}

	return emptyChar

}

func (r RenderEngine) createCanvas() [][]rune {
	canvas := make([][]rune, r.WindowSize[1])
	for py, x := range canvas {
		x = make([]rune, r.WindowSize[0])
		for px := range x {
			x[px] = r.getRune(px, py)
		}
		canvas[py] = x
	}
	return canvas
}

func (r RenderEngine) Render() string {
	s := ""
	for _, row := range r.createCanvas() {
		for _, cell := range row {
			s += string(cell)
		}
		s += "\n"
	}
	return s
}

