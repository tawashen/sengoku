package main

import (
	"fmt"
	"sort"
	"strings"
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
