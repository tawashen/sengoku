package main

func InitializeGenerals() map[string]*General {
	generals := make(map[string]*General)

	data := []General{
		// --- 武将１.jpg ---
		{Name: "龍造寺隆信", Combat: 4, Politics: 3, Prestige: 2, Stipend: 2, Loyalty: 0},
		{Name: "武田信玄", Combat: 4, Politics: 5, Prestige: 5, Stipend: 5, Loyalty: 0, PlusLoyalty: -1},
		{Name: "徳川家康", Combat: 4, Politics: 3, Prestige: 3, Stipend: 4, Loyalty: 1},
		{Name: "羽柴秀吉", Combat: 3, Politics: 5, Prestige: 3, Stipend: 2, Loyalty: 2},
		{Name: "武田勝頼", Combat: 4, Politics: 0, Prestige: 3, Stipend: 4, Loyalty: 0},
		{Name: "織田信長", Combat: 3, Politics: 4, Prestige: 5, Stipend: 5, Loyalty: 0, PlusLoyalty: -1},
		{Name: "島津義久", Combat: 4, Politics: 4, Prestige: 4, Stipend: 4, Loyalty: 0, PlusLoyalty: -1},
		{Name: "山中鹿之介", Combat: 2, Politics: 0, Prestige: 1, Stipend: 1, Loyalty: 3},
		{Name: "山本勘助", Combat: 0, PlusCombat: 1, Politics: 0, Prestige: 0, Stipend: 1, Loyalty: 1},
		{Name: "上杉謙信", Combat: 5, Politics: 0, Prestige: 5, Stipend: 2, Loyalty: 1},
		{Name: "大内義隆", Combat: 2, Politics: 2, Prestige: 3, Stipend: 4, Loyalty: 0},
		{Name: "大友宗麟", Combat: 3, Politics: 2, Prestige: 3, Stipend: 5, Loyalty: 0, PlusLoyalty: -2},
		{Name: "松永久秀", Combat: 3, Politics: 4, Prestige: 1, Stipend: 1, Loyalty: 0, PlusLoyalty: -3},
		{Name: "滝川一益", Combat: 4, Politics: 2, Prestige: 3, Stipend: 2, Loyalty: 0},
		{Name: "吉川元春", Combat: 4, Politics: 2, Prestige: 3, Stipend: 3, Loyalty: 1},
		{Name: "斎藤道三", Combat: 4, Politics: 4, Prestige: 2, Stipend: 1, Loyalty: 0, PlusLoyalty: -3},
		{Name: "今川義元", Combat: 1, Politics: 2, Prestige: 5, Stipend: 5, Loyalty: 0, PlusLoyalty: -1},
		{Name: "柴田勝家", Combat: 3, Politics: 1, Prestige: 3, Stipend: 3, Loyalty: 2},
		{Name: "大友宗麟_2", Combat: 3, Politics: 2, Prestige: 3, Stipend: 5, Loyalty: 0, PlusLoyalty: -2},
		{Name: "立花道雪", Combat: 5, Politics: 4, Prestige: 3, Stipend: 3, Loyalty: 3},

		// --- 武将２.jpg ---
		{Name: "毛利元就", Combat: 3, Politics: 4, Prestige: 3, Stipend: 2, Loyalty: 0, PlusLoyalty: -3},
		{Name: "宇喜多直家", Combat: 3, Politics: 3, Prestige: 2, Stipend: 2, Loyalty: 0, PlusLoyalty: -2},
		{Name: "朝倉義景", Combat: 0, Politics: 0, Prestige: 3, Stipend: 5, Loyalty: 0, PlusLoyalty: -1},
		{Name: "北条氏康", Combat: 4, Politics: 4, Prestige: 4, Stipend: 5, Loyalty: 0},
		{Name: "小早川隆景", Combat: 3, Politics: 4, Prestige: 4, Stipend: 4, Loyalty: 1},
		{Name: "毛利輝元", Combat: 2, Politics: 2, Prestige: 4, Stipend: 4, Loyalty: 1},
		{Name: "竹中半兵衛", Combat: 0, PlusCombat: 1, Politics: 2, Prestige: 1, Stipend: 1, Loyalty: 2},
		{Name: "伊達政宗", Combat: 4, Politics: 3, Prestige: 4, Stipend: 3, Loyalty: 0, PlusLoyalty: -1},
		{Name: "浅井長政", Combat: 4, Politics: 2, Prestige: 2, Stipend: 3, Loyalty: 0},
		{Name: "山県昌景", Combat: 4, Politics: 3, Prestige: 4, Stipend: 3, Loyalty: 2},
		{Name: "長宗我部元親", Combat: 4, Politics: 3, Prestige: 3, Stipend: 2, Loyalty: 0},
		{Name: "黒田孝高", Combat: 0, PlusCombat: 1, Politics: 3, Prestige: 1, Stipend: 2, Loyalty: 0},
		{Name: "明智光秀", Combat: 5, Politics: 4, Prestige: 3, Stipend: 5, Loyalty: 0, PlusLoyalty: -2},
		{Name: "足利義昭", Combat: 0, Politics: 0, Prestige: 0, PlusPrestige: 2, Stipend: 8, Loyalty: 0, PlusLoyalty: -2},
		{Name: "筒井順慶", Combat: 2, Politics: 3, Prestige: 3, Stipend: 2, Loyalty: 0, PlusLoyalty: 3}, // ±3は補正として扱う
		{Name: "馬場信春", Combat: 4, Politics: 3, Prestige: 4, Stipend: 3, Loyalty: 2},
		{Name: "北条早雲", Combat: 3, Politics: 4, Prestige: 2, Stipend: 1, Loyalty: 0, PlusLoyalty: -2},
	}

	for _, g := range data {
		g.ID = g.Name
		generals[g.ID] = &g
	}

	return generals
}
