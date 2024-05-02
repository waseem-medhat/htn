package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model int
type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		m--
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("\n\n     Hi. This program will exit in %d seconds...", m)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(model(5), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	req, err := http.NewRequest(http.MethodGet, "https://wipdev.netlify.app", nil)
	if err != nil {
		return fmt.Errorf("error forming the request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending the request: %v", err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println(string(resBody))

	return err
}
