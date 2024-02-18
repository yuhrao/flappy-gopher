package renderer

const (
  MoveMsg EngineMessage = iota
  GravityMsg EngineMessage = iota
)

type EngineMessage int
