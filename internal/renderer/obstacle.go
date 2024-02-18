package renderer

type Obstacle struct {
	top    int
	bottom int
	px     int
	width  int
	gap    int
}

func NewObstacle(px, gap, height, width int) *Obstacle {
  halfGap := int(gap / 2)
	top := height - halfGap
	bottom := height + halfGap
	return &Obstacle{top, bottom, px, gap, width}
}

func (o *Obstacle) IntersectingY(y int) bool {
	return y <= o.top || y >= o.bottom
}

func (o *Obstacle) IntersectingX(x int) bool {
	return x <= o.px+o.width && x >= o.px
}

func (o *Obstacle) Move(length int) {
  o.px = o.px - 1
}
