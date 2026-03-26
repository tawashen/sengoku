package main

import (
	"fmt"
	"math/rand/v2"

	//"sort"
	//"strings"

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
