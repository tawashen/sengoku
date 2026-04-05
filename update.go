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

	case "調略結果判定フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter": //結果を表示してるだけなので次のフェイズへ
				m.gameState.Phase = "メッセージ表示フェイズ"
				m.gameState.Message = "調略結果を発表します"
				m.gameState.PhaseStorage = "徴税フェイズ"
				return m, nil
			}
		}
	case "調略フェイズ":
		player := m.gameState.Players[m.gameState.Order[m.gameState.PlayerCounter]] //現在のプレイヤー

		// 安全装置：カーソルが自分の領国の数を超えていたら安全な値に直す
		if len(player.SchemeProvinces) > 0 && m.cursor >= len(player.SchemeProvinces) {
			m.cursor = len(player.SchemeProvinces) - 1
		} else if len(player.SchemeProvinces) == 0 {
			// 万が一対象が0の場合はスキップ
			m.gameState.PlayerCounter++
			if m.gameState.PlayerCounter >= len(m.gameState.Players) {
				m.gameState.PlayerCounter = 0
				m.gameState.Phase = "メッセージ表示フェイズ"
				m.gameState.Message = "調略結果を発表します"
				m.gameState.PhaseStorage = "徴税フェイズ"
			} else {
				m.InitSchemeProvinces(m.gameState.PlayerCounter)
			}
			return m, nil
		}

		currentProvince := player.SchemeProvinces[m.cursor] //カーソルのある領国
		currentCounter := m.gameState.PlayerCounter   //現在のプレイヤーのカウンター
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				//資金バラマキ終了、次のプレイヤーへ
				m.gameState.PlayerCounter++
				//最後のプレイヤーだった場合には次のフェイズへ
				if m.gameState.PlayerCounter >= len(m.gameState.Players) {
					m.gameState.PlayerCounter = 0
					m.gameState.Phase = "メッセージ表示フェイズ"
					m.gameState.Message = "調略結果を発表します"
					m.gameState.PhaseStorage = "徴税フェイズ"
				} else {
					m.InitSchemeProvinces(m.gameState.PlayerCounter)
				}
				m.cursor = 0
				//カーソルの領国に資金を割り振り、数値入力
			case "right":
				if player.Gold == 0 {
					return m, nil
				}
				player.Gold--
				currentProvince.Bids[currentCounter]++
				return m, nil
			case "left":
				if currentProvince.Bids[currentCounter] == 0 {
					return m, nil
				}
				player.Gold++
				currentProvince.Bids[currentCounter]--
				return m, nil
			}
		}

	case "徴税フェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":

				if m.gameState.Players[m.gameState.PlayerCounter].EventC.Tax {
					m.ExecuteCard(m.gameState.Players[m.gameState.PlayerCounter].EventC)
					return m, nil
				}

				// //現在のプレイヤーのイベントカードが「飢饉」なら飢饉感染チェックフェイズへ
				// if m.gameState.Players[m.gameState.PlayerCounter].EventC.Name == "飢饉" {
				// 	m.gameState.Phase = "メッセージ表示フェイズ"
				// 	m.gameState.Message = "飢饉カード！飢饉が発生しました。隣接国に感染する可能性があります\n"
				// 	m.gameState.PhaseStorage = "飢饉感染チェックフェイズ"
				// 	return m, nil
				// }

				// //現在のプレイヤーのカードのが「国内の不穏」なら徴税選択フェイズへ
				// if m.gameState.Players[m.gameState.PlayerCounter].EventC.Name == "国内の不穏" {
				// 	m.gameState.Phase = "メッセージ表示フェイズ"
				// 	m.gameState.Message = "国内の不穏カード！徴税を行えばすべての支配国が不穏状態となります。それでも徴税しますか？(Y/N)\n"
				// 	m.gameState.PhaseStorage = "徴税選択フェイズ"
				// 	return m, nil
				// }
			}
			return m, nil
		}
	case "飢饉感染チェックフェイズ":
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":

				//飢饉発生国スライス作成
				starvingProvinces := []*Province{}
				for _, p := range m.gameState.Provinces {
					if p.Starving {
						starvingProvinces = append(starvingProvinces, p)
					}
				}

				//新規感染国リスト
				newlyInfectedProvinces := ""

				//飢饉発生国の隣接国をリストアップして飢饉チェック
				for _, sp := range starvingProvinces {
					for _, p := range sp.Neighbors {
						if !m.gameState.Provinces[p].StarvingChecked {
							if rand.IntN(6)+1 > 4 {
								newlyInfectedProvinces += m.gameState.Provinces[p].Name + ","
								m.gameState.Provinces[p].Starving = true
								m.gameState.Provinces[p].StarvingChecked = true
							} else {
								m.gameState.Provinces[p].StarvingChecked = true
							}
						}
					}
				}
				//新規感染国がなければ終了
				if newlyInfectedProvinces == "" {
					m.gameState.Phase = "メッセージ表示フェイズ"
					m.gameState.Message = "飢饉は感染しませんでした"
					m.gameState.PhaseStorage = "徴税フェイズ"
					//現在のプレイヤーのイベントカードを消去
					m.gameState.Players[m.gameState.PlayerCounter].EventC = Card{}
					return m, nil
				}
				//新規感染国があれば
				m.gameState.Phase = "メッセージ表示フェイズ"
				m.gameState.Message = "飢饉が" + newlyInfectedProvinces + "に感染しました"
				m.gameState.PhaseStorage = "飢饉感染チェックフェイズ"
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

				// もしこれから「調略フェイズ」に入るなら、最初のプレイヤーの対象国リストを初期化する
				if m.gameState.Phase == "調略フェイズ" {
					m.InitSchemeProvinces(m.gameState.PlayerCounter)
				}

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

// InitSchemeProvinces は、大名が調略フェイズでアクセスできる領国リスト（自分の領国＋隣接国）を生成して保存します
func (m *model) InitSchemeProvinces(playerIdx int) {
	if len(m.gameState.Order) == 0 {
		return
	}
	player := m.gameState.Players[m.gameState.Order[playerIdx]]
	seen := make(map[string]bool)
	player.SchemeProvinces = []*Province{}

	// 1. 自分の領国を追加
	for _, p := range player.Provinces {
		if !seen[p.Name] {
			seen[p.Name] = true
			player.SchemeProvinces = append(player.SchemeProvinces, p)
		}
	}

	// 2. 隣接国を追加
	for _, p := range player.Provinces {
		for _, nID := range p.Neighbors {
			if neighbor, ok := m.gameState.Provinces[nID]; ok {
				if !seen[neighbor.Name] {
					seen[neighbor.Name] = true
					player.SchemeProvinces = append(player.SchemeProvinces, neighbor)
				}
			}
		}
	}
}
