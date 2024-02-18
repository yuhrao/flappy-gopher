package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yuhrao/flappy-gopher/internal/renderer"
)

type model struct {
	engine *renderer.Engine
}

func initialModel() model {
	wSize := [2]int{130, 50}

	return model{
		engine: renderer.NewEngine(wSize),
	}
}


func movementCmd() tea.Cmd {
  return tea.Tick(500 * time.Millisecond, func(time.Time) tea.Msg {
    return renderer.MoveMsg
  })
}

func gravityCmd() tea.Cmd {
  return tea.Tick(500 * time.Millisecond, func(time.Time) tea.Msg {
    return renderer.GravityMsg
  })
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return tea.Batch(movementCmd(), gravityCmd())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

    // These keys should exit the program.
    case "ctrl+c", "q":
      return m, tea.Quit

    case " ":
      m.engine.Jump()
    }

	case renderer.EngineMessage:
		switch msg {
		case renderer.MoveMsg:
      m.engine.MoveObstacles()
      return m, movementCmd()

    case renderer.GravityMsg:
      m.engine.ApplyGravity()
      return m, gravityCmd()

		}

	}

	return m, nil
}

func (m model) View() string {
	return m.engine.Render()
}

func main() {

	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithFPS(120))

	if _, err := p.Run(); err != nil {

		fmt.Println("Error running program: ", err)
		panic(err)

	}

}
