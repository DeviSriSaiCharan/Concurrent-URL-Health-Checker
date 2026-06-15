package loader

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
)

// model contains the program state and its core function (Init, Update, View)
type model struct {
	Spinner   spinner.Model
	Total     int
	Completed int
	Done      bool
	// receive only channal that receives notifications from the workers
	ProgressChan <-chan int
}

type ProgressMsg struct {
	completed int
}

// create a new bubble tea model with given total and progress channel (like a constructor)
func New(total int, progressC <-chan int) tea.Model {
	sp := spinner.New()
	sp.Spinner = spinner.Globe
	return model{
		Spinner:      sp,
		Total:        total,
		ProgressChan: progressC,
		Completed:    0,
	}
}

// When bubble tea starts
// -> it starts the spinner
// -> wait for the progress channal to receive the message
// -> both run concurrently (spinner and waiting for progress channal)
func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		waitForProgress(m.ProgressChan),
	)
}

func waitForProgress(progressC <-chan int) tea.Cmd {
	return func() tea.Msg {
		// waits until the worker sends the message
		n := <-progressC
		return ProgressMsg{
			completed: n,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressMsg:
		m.Completed = msg.completed

		if m.Completed >= m.Total {
			m.Done = true
			return m, tea.Quit
		}

		return m, waitForProgress(m.ProgressChan)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	percent := (float64(m.Completed) / float64(m.Total)) * 100

	s := fmt.Sprintf("\n%s Checking URLs...\nCompleted: %d / %d (%.1f%%)\n",
		m.Spinner.View(), m.Completed, m.Total, percent)

	return s
}
