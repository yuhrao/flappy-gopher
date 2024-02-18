package renderer;

type Obstacle struct {
	topPos    [2]int
	bottomPos [2]int
	gap       int
}

func NewObstacle(px, gap, height int) Obstacle {
	topPos := [2]int{px, height - int(gap/2)}
	bottomPos := [2]int{px, height + int(gap/2)}
	return Obstacle{topPos, bottomPos, gap}
}

func (o Obstacle) IntersectingY(y int) bool {
	return y <= o.topPos[1] || y >= o.bottomPos[1]
}

func (o Obstacle) IntersectingX(x int) bool {
	return x <= o.topPos[0]+obstacleWidth && x >= o.topPos[0]
}
