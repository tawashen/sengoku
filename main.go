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
		Year:           1560,
		Phase:          "順番決定フェイズ",
		Provinces:      InitializeProvinces(),
		Cards:          InitializeCards(),
		Order:          []int{0, 1, 2},
		CardCount:      0,
		GeneralCounter: 0,
		PlayerCounter:  0,
		GeneralsList:   []*General{},
		Message:        "",
	}

	m := model{
		gameState: gs,
		cursor:    0,
	}

	gs.Generals = m.InitializeGenerals()

	// プレイヤーの初期化 (武将データからの参照)
	nobunaga := gs.Generals["織田信長"]
	nobunaga.Stipend = 0
	nobunaga.ProvinceID = "尾張"
	nobunaga.OwnerID = "織田"

	// 大名（プレイヤー）としての拡張フィールド
	nobunaga.Gold = 100
	nobunaga.Clan = "織田"
	nobunaga.IsAI = false
	nobunaga.Vassals = []*General{}
	nobunaga.Vassals = append(nobunaga.Vassals, gs.Generals["柴田勝家"])
	nobunaga.Vassals = append(nobunaga.Vassals, gs.Generals["滝川一益"])
	nobunaga.Vassals = append(nobunaga.Vassals, gs.Generals["明智光秀"])
	nobunaga.Provinces = []*Province{}
	nobunaga.Provinces = append(nobunaga.Provinces, gs.Provinces["尾張"])
	nobunaga.Provinces = append(nobunaga.Provinces, gs.Provinces["美濃"])
	nobunaga.Power = 0 //後で国力計算メソッド必要
	nobunaga.EventC = Card{}
	nobunaga.SecretC = []Card{}

	//尾張の初期化
	owari := gs.Provinces["尾張"]
	owari.OwnerID = "織田"
	owari.Complete = true
	owari.Castles = []*Castle{}
	owari.Castles = append(owari.Castles, &Castle{
		Ruler: "織田",
		Power: 5,
	})
	owari.Soldiers = 1
	owari.Restless = false
	owari.HasUprising = false
	owari.Starving = false
	owari.Christian = false
	owari.TradePort = false
	owari.GoldMine = false
	owari.Ikko = false
	owari.Honganji = false
	owari.Region = 1
	owari.Neighbors = []string{"美濃", "三河", "伊勢"}

	shingen := gs.Generals["武田信玄"]
	shingen.Stipend = 0
	shingen.ProvinceID = "甲斐"
	shingen.OwnerID = "武田"
	shingen.Gold = 100
	shingen.Clan = "武田"
	shingen.IsAI = false
	shingen.Vassals = []*General{}
	if gs.Generals["武田勝頼"] != nil {
		shingen.Vassals = append(shingen.Vassals, gs.Generals["武田勝頼"])
	}
	if gs.Generals["武田信廉"] != nil {
		shingen.Vassals = append(shingen.Vassals, gs.Generals["武田信廉"])
	}
	shingen.Provinces = []*Province{}
	shingen.Provinces = append(shingen.Provinces, gs.Provinces["甲斐"])
	shingen.Provinces = append(shingen.Provinces, gs.Provinces["信濃(北)"])
	shingen.Provinces = append(shingen.Provinces, gs.Provinces["信濃(南)"])
	shingen.Power = 0 //後で国力計算メソッド必要
	shingen.EventC = Card{}
	shingen.SecretC = []Card{}

	kenshin := gs.Generals["上杉謙信"]
	kenshin.Stipend = 0
	kenshin.ProvinceID = "越後"
	kenshin.OwnerID = "上杉"
	kenshin.Gold = 100
	kenshin.Clan = "上杉"
	kenshin.IsAI = false
	kenshin.Vassals = []*General{}
	if gs.Generals["上杉景勝"] != nil {
		kenshin.Vassals = append(kenshin.Vassals, gs.Generals["上杉景勝"])
	}
	if gs.Generals["上杉景虎"] != nil {
		kenshin.Vassals = append(kenshin.Vassals, gs.Generals["上杉景虎"])
	}
	kenshin.Provinces = []*Province{}
	kenshin.Provinces = append(kenshin.Provinces, gs.Provinces["越後"])
	kenshin.Provinces = append(kenshin.Provinces, gs.Provinces["越中"])
	kenshin.Power = 0 //後で国力計算メソッド必要
	kenshin.EventC = Card{}
	kenshin.SecretC = []Card{}

	gs.Players = []*General{
		nobunaga,
		shingen,
		kenshin,
	}

	return model{
		gameState: gs,
		cursor:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

		case "s": //ステータス表示
			m.gameState.PhaseStorage = m.gameState.Phase
			m.gameState.Phase = "ステータス表示フェイズ"

		case "space": //messageを消すため
			m.gameState.Message = ""
		}
	}

	switch m.gameState.Phase {
	case "徴税選択フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "y":
				m.gameState.Phase = "メッセージ表示フェイズ"
				m.gameState.Message = "徴税を行います"
				for _, p := range m.gameState.Players[m.gameState.PlayerCounter].Provinces {
					p.Restless = true
				}
				m.gameState.Players[m.gameState.PlayerCounter].EventC = Card{}
				return m, nil
			case "n":
				m.gameState.Phase = "メッセージ表示フェイズ"
				m.gameState.Message = "徴税を行いません"
				m.gameState.Players[m.gameState.PlayerCounter].EventC = Card{}
				m.gameState.PlayerCounter++
				return m, nil
			}
			return m, nil
		}

	case "順番決定フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = "武将登用フェイズ"
				m.gameState.Order = MyShuffleInt(m.gameState.Order)
				m.DistributeCards()
				return m, nil
			}
			return m, nil
		}

	case "ステータス表示フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = m.gameState.PhaseStorage
				m.gameState.Message = ""
				return m, nil
			}
			return m, nil
		}

	case "メッセージ表示フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = m.gameState.PhaseStorage
				m.gameState.Message = ""

				return m, nil
			}
			return m, nil
		}
	case "武将登用フェイズ":
		// 安全装置：もしGeneralsListが空ならスキップ
		if len(m.gameState.GeneralsList) == 0 {
			m.gameState.Phase = "吉凶札配布フェイズ"
			return m, nil
		}

		// (Game initialization/Round start should shuffle GeneralsList if needed)
		if m.gameState.GeneralCounter >= len(m.gameState.GeneralsList) {
			m.gameState.Phase = "吉凶札配布フェイズ"
			return m, nil
		}

		// 全員のターンが終わったら次のフェイズへ
		if m.gameState.PlayerCounter >= len(m.gameState.Players) {
			m.gameState.PlayerCounter = 0
			m.gameState.Phase = "吉凶札配布フェイズ"
			return m, nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "n":
				m.gameState.Message = fmt.Sprintf("%sは不採用となりました", m.gameState.GeneralsList[m.gameState.GeneralCounter].Name)
				m.gameState.GeneralCounter++
				m.gameState.PlayerCounter++
				m.gameState.PhaseStorage = m.gameState.Phase
				m.gameState.Phase = "メッセージ表示フェイズ"

				// 全員終わったかチェック
				if m.gameState.PlayerCounter >= len(m.gameState.Order) {
					m.gameState.PlayerCounter = 0
					m.gameState.Phase = "メッセージ表示フェイズ"
					m.gameState.PhaseStorage = "吉凶札配布フェイズ"
				}
				return m, nil

			case "y":
				m.gameState.Message = fmt.Sprintf("%sは採用となりました", m.gameState.GeneralsList[m.gameState.GeneralCounter].Name)
				currentPlayerIdx := m.gameState.Order[m.gameState.PlayerCounter]
				// 配下に加える
				daimyo := m.gameState.Players[currentPlayerIdx]
				daimyo.Vassals = append(daimyo.Vassals, m.gameState.GeneralsList[m.gameState.GeneralCounter])

				// 採用でも不採用でも次の武将へ、そして次のプレイヤーへ
				m.gameState.GeneralCounter++
				m.gameState.PlayerCounter++
				m.gameState.PhaseStorage = m.gameState.Phase
				m.gameState.Phase = "メッセージ表示フェイズ"

				// 全員終わったかチェック
				if m.gameState.PlayerCounter >= len(m.gameState.Order) {
					m.gameState.PlayerCounter = 0
					m.gameState.Phase = "メッセージ表示フェイズ"
					m.gameState.PhaseStorage = "吉凶札配布フェイズ"
				}
				return m, nil
			}
		}
		return m, nil

	case "吉凶札配布フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = "吉凶札実行フェイズ"
				m.gameState.PhaseStorage = "吉凶札配布フェイズ"
				return m, nil
			}
			return m, nil
		}

	case "吉凶札実行フェイズ":

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":

				//吉凶札を持っていない場合
				if m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]].EventC == (Card{}) {
					m.gameState.Message = fmt.Sprintf("%sは事件札を持っていません", m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]].Name)
					m.gameState.Phase = "メッセージ表示フェイズ"
					m.gameState.PhaseStorage = "吉凶札実行フェイズ"
				} else {
					card := m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]].EventC
					//吉凶札が徴税札の場合
					if card.Tax {
						m.gameState.Message = fmt.Sprintf("%sは徴税フェイズで使います", card.Name)
						m.gameState.Phase = "メッセージ表示フェイズ"
						m.gameState.PhaseStorage = "吉凶札実行フェイズ"
					} else {
						//サイコロを振る
						if card.Dice != nil {
							m.gameState.DiceResult = rand.IntN(6) + 1
						}
						//カード実行（自動的に Phase = "メッセージ表示フェイズ", PhaseStorage = "吉凶札実行フェイズ" がセットされる想定）
						m.ExecuteCard(card)
					}
				}

				// 処理が終わったら次のプレイヤーへ
				m.gameState.PlayerCounter++

				// もし今終わったのが「最後のプレイヤー」だった場合、
				// メッセージ表示後に戻る先(PhaseStorage)を "調略フェイズ" に変更する！
				if m.gameState.PlayerCounter >= len(m.gameState.Order) {
					m.gameState.PhaseStorage = "調略フェイズ"
					m.gameState.PlayerCounter = 0
				}

				return m, nil
			}
			return m, nil
		}
		return m, nil
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

	case "吉凶札実行フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("吉凶札実行フェイズ:\n\n")
		s.WriteString(m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]].Name)
		s.WriteString("のターンです Please Enter\n")
		s.WriteString(m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]].EventC.Name)
		s.WriteString("\n\n")
		s.WriteString(m.gameState.Message)
		s.WriteString("\n\n")

	case "徴税選択フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("徴税選択フェイズ:\n\n")
		s.WriteString(m.gameState.Message)
		s.WriteString("\n\n")

	case "順番決定フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("順番決定フェイズ:\n\n")
		s.WriteString("(enter: 次のフェイズへ, c: 国のつながり確認)\n")

	case "ステータス表示フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("ステータス表示フェイズ:\n\n")
		s.WriteString("(enter: 元のフェイズへ)\n")

		counter := m.gameState.PlayerCounter
		if counter >= len(m.gameState.Order) {
			counter = 0
		}
		p := m.gameState.Players[m.gameState.Order[counter]]

		provinces := ""
		for _, province := range p.Provinces {
			provinces += province.Name + ", "
		}
		s.WriteString(fmt.Sprintf("大名：%s　領地：%s　\n", p.Name, provinces))
		s.WriteString("家臣：\n")
		for _, general := range p.Vassals {
			s.WriteString(fmt.Sprintf("%s　\n", general.Name))
		}
		s.WriteString("領地の詳細：\n")
		for _, province := range p.Provinces {
			castles := ""
			for _, castle := range province.Castles {
				castles += castle.Ruler + ", "
			}
			s.WriteString(fmt.Sprintf("%s: 国力：%d　兵士：%d　城塞：%s　\n", province.Name, province.Kokuryoku, province.Soldiers, castles))
		}
		s.WriteString("城塞の詳細：\n")
		for _, province := range p.Provinces {
			for _, castle := range province.Castles {
				s.WriteString(fmt.Sprintf("%s: 支配者：%s　\n", castle.Ruler, castle.Ruler))
			}
		}
		s.WriteString(m.gameState.Message)

	case "メッセージ表示フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("メッセージ表示フェイズ:\n\n")
		s.WriteString("(enter: 元のフェイズへ)\n")
		s.WriteString(m.gameState.Message)

	case "武将登用フェイズ":
		s.WriteString("\n\n")
		s.WriteString("武将登用フェイズ:\n\n")

		if len(m.gameState.Order) > 0 && m.gameState.PlayerCounter < len(m.gameState.Order) {
			currentPlayerIdx := m.gameState.Order[m.gameState.PlayerCounter]
			s.WriteString(fmt.Sprintf("現在のプレイヤー: %s\n", m.gameState.Players[currentPlayerIdx].Name))
		}

		if m.gameState.GeneralCounter < len(m.gameState.GeneralsList) {
			s.WriteString(fmt.Sprintf("%s　この武将を採用しますか？(y/n)\n", m.gameState.GeneralsList[m.gameState.GeneralCounter].Name))
		}
		s.WriteString(m.gameState.Message)

	case "吉凶札配布フェイズ":
		s.WriteString("\n\n")
		//s.WriteString("吉凶札配布フェイズ:\n\n")
		s.WriteString("Enter　吉凶札実行\n")
		// for i, pIdx := range m.gameState.Order {
		// 	daimyo := m.gameState.Players[pIdx][0]
		// 	secretcards := ""
		// 	for _, card := range daimyo.SecretC {
		// 		secretcards += card.Name + ", "
		// 	}
		daimyo := m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]]
		secretcards := ""
		for _, card := range daimyo.SecretC {
			secretcards += card.Name + ", "
		}
		s.WriteString(fmt.Sprintf("%s (事件札: %s) (秘密札: %s)\n", daimyo.Name, daimyo.EventC.Name, secretcards))

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
	m := initialModel()
	p := tea.NewProgram(&m)
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
	for _, daimyo := range m.gameState.Players {
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
