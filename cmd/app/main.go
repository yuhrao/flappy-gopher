package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yuhrao/flappy-gopher/internal/renderer"
)

type model struct {
	engine renderer.RenderEngine
}

func initialModel() model {
	wSize := [2]int{100, 30}
	obstacles := []renderer.Obstacle{
		renderer.NewObstacle(19, 4, 15),
		renderer.NewObstacle(39, 2, 10),
	}

	return model{
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		engine: renderer.RenderEngine{
			WindowSize: wSize,
			BirdHeight: 12,
			Obstacles:  obstacles,
		},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
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

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// TODO: Check if we should re-render the view if nothing has changed

	return m.engine.Render()
}

func main() {

	fmt.Println("Initialising...")
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {

		fmt.Println("Error running program: ", err)
		panic(err)

	}

}
