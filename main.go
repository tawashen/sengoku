package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Main Model for Bubble Tea
type model struct {
	gameState *GameState
	cursor    int // Current selection in province list
}

func initialModel() model {
	// Sample data for visualization
	gs := &GameState{
		Year:  1560,
		Phase: "順番決定フェイズ",
		Provinces: map[string]*Province{
			"owari": {
				ID:        "owari",
				Name:      "尾張",
				Kokudaka:  10,
				Neighbors: []string{"mikawa", "mino", "ise"},
			},
			"mikawa": {
				ID:        "mikawa",
				Name:      "三河",
				Kokudaka:  5,
				Neighbors: []string{"owari", "totomi", "shinano"},
			},
			"mino": {
				ID:        "mino",
				Name:      "美濃",
				Kokudaka:  8,
				Neighbors: []string{"owari", "omi", "echizen", "shinano"},
			},
		},
		Players: []*Player{
			{ID: "nobunaga", Name: "織田信長", Clan: "織田"},
			{ID: "shingen", Name: "武田信玄", Clan: "武田"},
			{ID: "kenshin", Name: "上杉謙信", Clan: "上杉"},
		},
		Cards:    InitializeCards(),
		Generals: InitializeGenerals(),
		Order:    []int{0, 1, 2},
	}
	return model{
		gameState: gs,
		cursor:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
			// case "up", "k":
			// 	if m.cursor > 0 {
			// 		m.cursor--
			// 	}
			// case "down", "j":
			// 	if m.cursor < len(m.gameState.Provinces)-1 {
			// 		m.cursor++
			// 	}

		}
	}

	switch m.gameState.Phase {
	case "順番決定フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = "吉凶札配布フェイズ"
				m.gameState.Order = MyShuffleInt(m.gameState.Order)
				return m, nil
			}
			return m, nil
		}
	case "吉凶札配布フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = "戦闘フェイズ"
				return m, nil
			}
			return m, nil
		}
	}
	return m, nil
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	provinceStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true).
			PaddingLeft(0)
)

func (m model) View() string {
	s := strings.Builder{}

	// Header
	s.WriteString(titleStyle.Render(fmt.Sprintf("戦国大名 TUI - %d年 %sフェイズ", m.gameState.Year, m.gameState.Phase)))

	s.WriteString("\n\n")

	switch m.gameState.Phase {
	case "順番決定フェイズ":
		s.WriteString("\n\n")
		s.WriteString("順番決定フェイズ:\n\n")
		s.WriteString("\n(enter: 次のフェイズへ)\n")

	case "吉凶札配布フェイズ":
		s.WriteString("\n\n")
		s.WriteString("吉凶札配布フェイズ:\n\n")

	case "戦闘フェイズ":
		s.WriteString("\n\n")
		s.WriteString("戦闘フェイズ:\n\n")
	case "国のつながり確認フェイズ":
		s.WriteString("\n\n")
		s.WriteString("国のつながり確認フェイズ:\n\n")

		ids := []string{"owari", "mikawa", "mino"}
		for i, id := range ids {
			p := m.gameState.Provinces[id]
			cursor := "  "
			if m.cursor == i {
				cursor = "> "
			}

			neighborNames := []string{}
			for _, nid := range p.Neighbors {
				if neighbor, ok := m.gameState.Provinces[nid]; ok {
					neighborNames = append(neighborNames, neighbor.Name)
				} else {
					neighborNames = append(neighborNames, fmt.Sprintf("%s(未登録)", nid))
				}
			}

			line := fmt.Sprintf("%s%s (隣接: %s)", cursor, p.Name, strings.Join(neighborNames, ", "))
			if m.cursor == i {
				s.WriteString(selectedStyle.Render(line))
			} else {
				s.WriteString(provinceStyle.Render(line))
			}
			s.WriteString("\n")
		}
		s.WriteString("\n(↑↓: 選択, q: 終了)\n")
	}

	return s.String()
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}

}

func MyShuffleInt(list []int) []int {
	n := len(list)
	if n == 0 {
		return list
	}

	// シャッフル
	rand.Shuffle(n, func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	return list
}

func (m *model) ExecuteCard(c Card) {
	if effect, ok := EffectMap[c.Name]; ok {
		effect(m, c)
	}
}
