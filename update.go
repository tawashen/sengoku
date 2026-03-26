package main

import (
	"fmt"
	"math/rand/v2"

	tea "github.com/charmbracelet/bubbletea"
)

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

	case "飢饉感染チェックフェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.gameState.Phase = "徴税選択フェイズ"
				return m, nil
			}
			return m, nil
		}

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
