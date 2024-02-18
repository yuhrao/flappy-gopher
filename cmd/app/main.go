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
	wSize := [2]int{100, 30}

	return model{
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		engine: renderer.NewEngine(wSize),
	}
}

func gameCommands() tea.Cmd {
	return tea.Tick(100 * time.Millisecond, func(time.Time) tea.Msg {
		return renderer.MoveMsg
	})
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return gameCommands()
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
		}

	case renderer.EngineMessage:
		switch msg {
		case renderer.MoveMsg:
      m.engine.Move()
      return m, gameCommands()
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
  // var s string
  // for _, o := range m.engine.Obstacles {
  //   s += fmt.Sprint(*o)
  // }
  // return s

	return m.engine.Render()
}

func main() {

	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithFPS(120))

	if _, err := p.Run(); err != nil {

		fmt.Println("Error running program: ", err)
		panic(err)

	}

}
