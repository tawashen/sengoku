package main

scenaio := Scenario {
    Title: "四国統一",
    Member: 1,
    Provinces: InitializeProvinces(),
    Generals: InitializeGenerals(),
    Players: InitializePlayers(),
}

func InitializeScenario(s Scenario) GameState {
	orderN := for i := 0; i < s.Member; i++ {
		append(orderN, i)
	}
	orderN = MyShuffleInt(orderN)
	
	retrun GameState {
		Year: 1560,
		Phase: "順番決定フェイズ",
		Provinces: s.Provinces,
		Generals: s.Generals,
		Players: s.Players,
		Order: orderN,
		Cards: InitializeCards(),
		CardCount: 0,
	}

} 


//無名武将チップを読み取ってインスタンス化
//