package main

import "fmt"

// EffectMap はカード名と実際の処理内容を紐付けます。
// 全てのカード名が含まれていますので、中身を実装してください。
var EffectMap = map[string]func(m *model, c Card){
	"国内の不穏": func(m *model, c Card) {
		// TODO: 徴税フェイズに支配地域を一揆チェック状態にする
		m.gameState.Phase = "徴税選択フェイズ"
		m.gameState.PhaseStorage = "徴税フェイズ"
		m.gameState.Message = "不穏カードの影響により徴税を行えばすべての支配国が不穏状態となります。それでも徴税しますか？(Y/N)"

	},
	"一向一揆": func(m *model, c Card) {
		// ダイスを振り、指定地域に一向一揆を発生させる
		//ここで発生した一向一揆は隣接する国に波及する可能性がある、一向一揆波及フェイズを新たに作る
		provinceName := c.Dice.Result[m.gameState.DiceResult-1].(string)
		if p, ok := m.gameState.Provinces[provinceName]; ok {
			p.HasUprising = true
		}
		m.gameState.Phase = "メッセージ表示フェイズ"
		m.gameState.Message = fmt.Sprintf("%sに一向一揆が発生しました", provinceName)
		m.gameState.PhaseStorage = "吉凶札実行フェイズ"
	},
	"廃鉱": func(m *model, c Card) {
		// TODO: 金山マーカーを除去し、以後の収入を断つ
		provinces := ""
		player := m.gameState.Players[m.gameState.PlayerCounter]
		for _, p := range player.Provinces {
			if p.GoldMine {
				p.GoldMine = false
				provinces += p.Name + ", "
			}
		}
		m.gameState.Phase = "メッセージ表示フェイズ"
		m.gameState.Message = fmt.Sprintf("%sの金山が廃鉱となりました", provinces)
		m.gameState.PhaseStorage = "吉凶札実行フェイズ"
	},
	"裏切り": func(m *model, c Card) {
		// TODO: 俸禄最小の家臣を特定し、離脱させる
		player := m.gameState.Players[m.gameState.PlayerCounter]
		minStipend := 10000
		minIndex := -1
		for i, v := range player.Vassals {
			if v.Stipend < minStipend {
				minStipend = v.Stipend
				minIndex = i
			}
		}

		if minIndex != -1 {
			m.gameState.Phase = "メッセージ表示フェイズ"
			m.gameState.Message = fmt.Sprintf("家臣である%sが裏切りました", player.Vassals[minIndex].Name)
			m.gameState.PhaseStorage = "吉凶札実行フェイズ"

			player.Vassals = append(player.Vassals[:minIndex], player.Vassals[minIndex+1:]...)
		} else {
			m.gameState.Phase = "メッセージ表示フェイズ"
			m.gameState.Message = "家臣がいません"
			m.gameState.PhaseStorage = "吉凶札実行フェイズ"
		}

	},
	"大名死亡": func(m *model, c Card) {
		// TODO: 継承処理と全家臣の忠誠チェック
		player := m.gameState.Players[m.gameState.PlayerCounter]

		if len(player.Vassals) == 0 {
			m.gameState.Phase = "メッセージ表示フェイズ"
			m.gameState.Players[m.gameState.PlayerCounter] = m.gameState.GeneralsList[m.gameState.GeneralCounter]
			m.gameState.GeneralCounter++
			m.gameState.Message = fmt.Sprintf("家臣がいません　%sは滅亡しました。次期当主は%sです。", player.Name, m.gameState.Players[m.gameState.PlayerCounter].Name)
			m.gameState.PhaseStorage = "吉凶札実行フェイズ"
		}
	},
	"内応の露顕": func(m *model, c Card) {
		// TODO: 「城方の内応」を無効化し、強襲を強制する

	},
	"城方の内応": func(m *model, c Card) {
		// TODO: 城塞の除去、またはダイスによる戦力低下
	},
	"詭計の成功": func(m *model, c Card) {
		// TODO: 敵軍分割と戦闘ボーナスの付与
	},
	"伏兵": func(m *model, c Card) {
		// TODO: 防御側ボーナスと敵カードの無効化
	},
	"奇襲": func(m *model, c Card) {
		// TODO: 野戦強制とダイスボーナスの付与
	},
	"飢饉": func(m *model, c Card) {
		// TODO: 陸奥・出羽の混乱状態化と近隣への波及チェック
		//一向一揆と一緒で感染用フェイズを用意する必要あり
		for _, p := range m.gameState.Provinces {
			p.StarvingChecked = false
		}
		//飢饉チェック
		m.gameState.Provinces["陸奥(北)"].Starving = true
		m.gameState.Provinces["陸奥(北)"].StarvingChecked = true
		m.gameState.Provinces["陸奥(中)"].Starving = true
		m.gameState.Provinces["陸奥(中)"].StarvingChecked = true
		m.gameState.Provinces["陸奥(南)"].Starving = true
		m.gameState.Provinces["陸奥(南)"].StarvingChecked = true
		m.gameState.Provinces["出羽"].Starving = true
		m.gameState.Provinces["出羽"].StarvingChecked = true

		m.gameState.Phase = "メッセージ表示フェイズ"
		m.gameState.Message = "飢饉が発生しました"
		m.gameState.PhaseStorage = "飢饉感染チェックフェイズ"
	},
	"豊作": func(m *model, c Card) {
		// TODO: ダイスで指定された地域を豊作状態にする
	},
	"影武者": func(m *model, c Card) {
		// TODO: 死亡イベントを無効化する
	},
	"反対勢力の増大": func(m *model, c Card) {
		// TODO: 各領国で中立勢力の発生チェック
	},
	"凶作": func(m *model, c Card) {
		// TODO: ダイスで指定された地域を凶作状態にする
	},
	"南蛮人の渡来": func(m *model, c Card) {
		// TODO: ダイスに基づき南蛮貿易港マーカーを設置
	},
	"築城技術の進歩": func(m *model, c Card) {
		// TODO: 指定した城の戦力を5アップさせる
	},
	"九州での反対勢力の増大": func(m *model, c Card) {
		// TODO: 九州地方限定の中立勢力発生チェック
	},
	"土一揆との和睦": func(m *model, c Card) {
		// TODO: 土一揆の発生を防止する
	},
	"威信の上昇": func(m *model, c Card) {
		// TODO: 威信+1、金-20の処理
	},
	"隣国での混乱": func(m *model, c Card) {
		// TODO: 隣国の中立勢力発生と一部吸収処理
	},
	"キリシタンの流行": func(m *model, c Card) {
		// TODO: キリシタンマーカーの設置と波及
	},
	"他の勢力の調整": func(m *model, c Card) {
		// TODO: 自国または隣国の勢力5を吸収
	},
	"金山発見": func(m *model, c Card) {
		// TODO: 金山マーカー設置と永続収入追加
	},
	"国一揆との和睦": func(m *model, c Card) {
		// TODO: 国一揆の発生を防止する
	},
	"一向宗との和睦": func(m *model, c Card) {
		// TODO: 一向一揆の発生を防止する
	},
	"刺客": func(m *model, c Card) {
		// TODO: 指定した他家武将の暗殺チェック
	},
	"大義": func(m *model, c Card) {
		// TODO: 山城支配者への勝利得点付与
	},
	"関東での反対勢力の増大": func(m *model, c Card) {
		// TODO: 関東地方限定の中立勢力発生チェック
	},
	"陸奥・出羽での反対勢力の増大": func(m *model, c Card) {
		// TODO: 東北地方限定の中立勢力発生チェック
	},
	"敵方家臣の調略": func(m *model, c Card) {
		// TODO: 他家家臣への調略工作
	},
	"下剋上": func(m *model, c Card) {
		// TODO: 大名交代と全家臣の裏切りチェック、俸禄倍増
	},
	"富の浪費": func(m *model, c Card) {
		// TODO: 所持金半減の処理
	},
	"不穏工作": func(m *model, c Card) {
		// TODO: 指定地域を不穏状態にする工作
	},
	"家臣死亡": func(m *model, c Card) {
		// TODO: ダイスによる家臣死亡の決定
	},
	"密報": func(m *model, c Card) {
		// TODO: 他プレイヤーの手札を確認し戦闘カードなら破棄
	},
	"賢臣の諫言": func(m *model, c Card) {
		// TODO: 凶となるカードの無効化
	},
}
