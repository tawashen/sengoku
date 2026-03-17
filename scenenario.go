package main

var DefaultScenario = Scenario{
	Title:     "四国統一",
	Member:    1,
	// Provinces: InitializeProvinces(), // コメントアウト: 初期化時に呼ぶか別途設定
	// Generals:  InitializeGenerals(),  // コメントアウト: modelのメソッドとして呼ぶ必要あり
	// Players:   InitializePlayers(),   // コメントアウト: 未定義
}

func InitializeScenario(s Scenario) GameState {
	var orderN []int
	for i := 0; i < s.Member; i++ {
		orderN = append(orderN, i)
	}
	orderN = MyShuffleInt(orderN)

	return GameState{
		Year:      1560,
		Phase:     "順番決定フェイズ",
		Provinces: s.Provinces,
		Generals:  s.Generals,
		Players:   s.Players,
		Order:     orderN,
		Cards:     InitializeCards(),
		CardCount: 0,
	}
}

//無名武将チップを読み取ってインスタンス化
//
