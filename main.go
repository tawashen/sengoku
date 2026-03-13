package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
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
	gs := &GameState{
		Year:      1560,
		Phase:     "順番決定フェイズ",
		Provinces: InitializeProvinces(),
		Generals:  InitializeGenerals(),
		Cards:     InitializeCards(),
		Order:     []int{0, 1, 2},
		CardCount: 0,
	}

	// プレイヤーの初期化 (武将データからの参照)
	nobunaga := gs.Generals["織田信長"]
	nobunaga.Gold = 100
	nobunaga.Clan = "織田"
	nobunaga.IsAI = false

	shingen := gs.Generals["武田信玄"]
	shingen.Gold = 100
	shingen.Clan = "武田"
	shingen.IsAI = false

	kenshin := gs.Generals["上杉謙信"]
	kenshin.Gold = 100
	kenshin.Clan = "上杉"
	kenshin.IsAI = false

	gs.Players = [][]*General{
		{nobunaga},
		{shingen},
		{kenshin},
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
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.gameState.Provinces)-1 {
				m.cursor++
			}
		case "c": // 'c' key to switch to adjacency check phase
			m.gameState.Phase = "国のつながり確認フェイズ"

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
				m.DistributeCards()
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
		s.WriteString("(enter: 次のフェイズへ, c: 国のつながり確認)\n")

	case "吉凶札配布フェイズ":
		s.WriteString("\n\n")
		s.WriteString("吉凶札配布フェイズ:\n\n")
		s.WriteString("順番: \n")
		for i, pIdx := range m.gameState.Order {
			daimyo := m.gameState.Players[pIdx][0]
			secretcards := ""
			for _, card := range daimyo.SecretC {
				secretcards += card.Name + ", "
			}
			s.WriteString(fmt.Sprintf("%d. %s (事件札: %s) (秘密札: %s)\n", i+1, daimyo.Name, daimyo.EventC.Name, secretcards))
		}

	case "戦闘フェイズ":
		s.WriteString("\n\n")
		s.WriteString("戦闘フェイズ:\n\n")
	case "国のつながり確認フェイズ":
		s.WriteString("\n\n")
		s.WriteString("国のつながり確認フェイズ:\n\n")

		// Sort IDs for consistent display
		ids := make([]string, 0, len(m.gameState.Provinces))
		for id := range m.gameState.Provinces {
			ids = append(ids, id)
		}
		// Simple alphabetical sort for now (ideally by Region)
		sort.Strings(ids)

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

			line := fmt.Sprintf("%s%s [地域%d] (隣接: %s)", cursor, p.Name, p.Region, strings.Join(neighborNames, ", "))
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

func (m *model) DistributeCards() {
	for _, group := range m.gameState.Players {
		daimyo := group[0] // 大名に配る
		index := m.gameState.CardCount
		if m.gameState.Cards[index].Event {
			daimyo.EventC = m.gameState.Cards[index]
		} else {
			daimyo.SecretC = append(daimyo.SecretC, m.gameState.Cards[index])
		}
		if m.gameState.CardCount >= len(m.gameState.Cards)-1 {
			m.gameState.Cards = InitializeCards()
			m.gameState.CardCount = 0
		} else {
			m.gameState.CardCount++
		}
	}
}
