package main

func InitializeProvinces() map[string]*Province {
	provinces := map[string]*Province{
		// --- 地域 1 ---
		"陸奥(北)": {ID: "陸奥(北)", Name: "陸奥(北)", Kokuryoku: 6, Region: 1, Neighbors: []string{"陸奥(中)", "出羽"}},
		"陸奥(中)": {ID: "陸奥(中)", Name: "陸奥(中)", Kokuryoku: 12, Region: 1, Neighbors: []string{"陸奥(北)", "陸奥(南)", "出羽"}},
		"陸奥(南)": {ID: "陸奥(南)", Name: "陸奥(南)", Kokuryoku: 16, Region: 1, Neighbors: []string{"陸奥(中)", "下野", "常陸", "出羽", "越後"}},
		"出羽":    {ID: "出羽", Name: "出羽", Kokuryoku: 12, Region: 1, Neighbors: []string{"陸奥(北)", "陸奥(中)", "陸奥(南)", "越後"}},

		// --- 地域 2 ---
		"武蔵": {ID: "武蔵", Name: "武蔵", Kokuryoku: 16, Region: 2, Neighbors: []string{"相模", "上野", "下総", "甲斐"}},
		"相模": {ID: "相模", Name: "相模", Kokuryoku: 6, Region: 2, Neighbors: []string{"武蔵", "伊豆"}},
		"上総": {ID: "上総", Name: "上総", Kokuryoku: 10, Region: 2, Neighbors: []string{"安房", "下総"}},
		"安房": {ID: "安房", Name: "安房", Kokuryoku: 2, Region: 2, Neighbors: []string{"上総"}},
		"下総": {ID: "下総", Name: "下総", Kokuryoku: 10, Region: 2, Neighbors: []string{"常陸", "下野", "武蔵", "上総", "上野"}},
		"常陸": {ID: "常陸", Name: "常陸", Kokuryoku: 14, Region: 2, Neighbors: []string{"下野", "下総", "陸奥(南)"}},

		// --- 地域 3 ---
		"下野":    {ID: "下野", Name: "下野", Kokuryoku: 10, Region: 3, Neighbors: []string{"陸奥(南)", "上野", "常陸", "下総"}},
		"上野":    {ID: "上野", Name: "上野", Kokuryoku: 12, Region: 3, Neighbors: []string{"下野", "武蔵", "信濃(北)", "下総", "越後"}},
		"信濃(北)": {ID: "信濃(北)", Name: "信濃(北)", Kokuryoku: 8, Region: 3, Neighbors: []string{"上野", "越後", "信濃(南)", "甲斐"}},
		"信濃(南)": {ID: "信濃(南)", Name: "信濃(南)", Kokuryoku: 6, Region: 3, Neighbors: []string{"信濃(北)", "甲斐", "三河", "尾張", "美濃", "遠江"}},
		"甲斐":    {ID: "甲斐", Name: "甲斐", Kokuryoku: 6, Region: 3, Neighbors: []string{"信濃(南)", "駿河", "武蔵", "信濃(北)"}, GoldMine: true},

		// --- 地域 4 ---
		"越後": {ID: "越後", Name: "越後", Kokuryoku: 14, Region: 4, Neighbors: []string{"信濃(北)", "信濃(南)", "出羽", "越中", "上野", "陸奥(南)"}, GoldMine: true},
		"越中": {ID: "越中", Name: "越中", Kokuryoku: 12, Region: 4, Neighbors: []string{"能登", "加賀", "越後", "飛騨"}, Ikko: true},
		"能登": {ID: "能登", Name: "能登", Kokuryoku: 6, Region: 4, Neighbors: []string{"越中", "加賀"}},
		"加賀": {ID: "加賀", Name: "加賀", Kokuryoku: 10, Region: 4, Neighbors: []string{"越中", "能登", "越前"}, Ikko: true},
		"越前": {ID: "越前", Name: "越前", Kokuryoku: 14, Region: 4, Neighbors: []string{"加賀", "若狭", "近江(北)"}, Ikko: true},
		"若狭": {ID: "若狭", Name: "若狭", Kokuryoku: 2, Region: 4, Neighbors: []string{"越前", "近江(北)", "丹後", "丹波"}},

		// --- 地域 5 ---
		"伊豆": {ID: "伊豆", Name: "伊豆", Kokuryoku: 2, Region: 5, Neighbors: []string{"相模", "駿河"}},
		"駿河": {ID: "駿河", Name: "駿河", Kokuryoku: 6, Region: 5, Neighbors: []string{"伊豆", "甲斐", "遠江"}, GoldMine: true},
		"遠江": {ID: "遠江", Name: "遠江", Kokuryoku: 6, Region: 5, Neighbors: []string{"駿河", "三河"}},
		"三河": {ID: "三河", Name: "三河", Kokuryoku: 8, Region: 5, Neighbors: []string{"遠江", "信濃(南)", "尾張"}, Ikko: true},
		"尾張": {ID: "尾張", Name: "尾張", Kokuryoku: 14, Region: 5, Neighbors: []string{"三河", "美濃", "伊勢", "近江(南)"}},
		"伊勢": {ID: "伊勢", Name: "伊勢", Kokuryoku: 18, Region: 5, Neighbors: []string{"尾張", "美濃", "近江(南)", "伊賀", "大和", "紀伊"}, Ikko: true, Honganji: true},

		// --- 地域 6 ---
		"飛騨":   {ID: "飛騨", Name: "飛騨", Kokuryoku: 2, Region: 6, Neighbors: []string{"越中", "美濃"}},
		"美濃":   {ID: "美濃", Name: "美濃", Kokuryoku: 14, Region: 6, Neighbors: []string{"尾張", "信濃(南)", "近江(北)", "伊勢", "飛騨", "近江(南)"}},
		"近江(北)": {ID: "近江(北)", Name: "近江(北)", Kokuryoku: 10, Region: 6, Neighbors: []string{"美濃", "近江(南)", "若狭", "越前"}},
		"近江(南)": {ID: "近江(南)", Name: "近江(南)", Kokuryoku: 14, Region: 6, Neighbors: []string{"近江(北)", "山城", "伊賀", "伊勢", "美濃"}, Ikko: true},
		"伊賀":   {ID: "伊賀", Name: "伊賀", Kokuryoku: 2, Region: 6, Neighbors: []string{"伊勢", "近江(南)", "大和", "山城"}},

		// --- 地域 7 ---
		"大和": {ID: "大和", Name: "大和", Kokuryoku: 12, Region: 7, Neighbors: []string{"伊勢", "伊賀", "紀伊", "河内", "山城"}},
		"山城": {ID: "山城", Name: "山城", Kokuryoku: 14, Region: 7, Neighbors: []string{"近江(南)", "摂津", "河内", "丹波", "大和", "伊賀"}},
		"摂津": {ID: "摂津", Name: "摂津", Kokuryoku: 10, Region: 7, Neighbors: []string{"山城", "和泉", "河内", "播磨", "丹波"}, Honganji: true, Ikko: true},
		"河内": {ID: "河内", Name: "河内", Kokuryoku: 6, Region: 7, Neighbors: []string{"山城", "大和", "摂津", "和泉", "紀伊"}},
		"和泉": {ID: "和泉", Name: "和泉", Kokuryoku: 2, Region: 7, Neighbors: []string{"摂津", "河内", "紀伊"}, TradePort: true},
		"播磨": {ID: "播磨", Name: "播磨", Kokuryoku: 12, Region: 7, Neighbors: []string{"摂津", "丹波", "備前", "美作", "因幡", "但馬"}},
		"周防": {ID: "周防", Name: "周防", Kokuryoku: 6, Region: 7, Neighbors: []string{"安芸", "長門", "石見"}},
		"丹波": {ID: "丹波", Name: "丹波", Kokuryoku: 6, Region: 7, Neighbors: []string{"山城", "摂津", "播磨", "丹後", "但馬", "若狭"}},

		// --- 地域 8 ---
		"長門": {ID: "長門", Name: "長門", Kokuryoku: 4, Region: 8, Neighbors: []string{"周防", "石見", "豊前"}},
		"丹後": {ID: "丹後", Name: "丹後", Kokuryoku: 4, Region: 8, Neighbors: []string{"若狭", "但馬", "丹波"}},
		"但馬": {ID: "但馬", Name: "但馬", Kokuryoku: 4, Region: 8, Neighbors: []string{"丹後", "因幡", "丹波", "播磨"}, GoldMine: true},
		"因幡": {ID: "因幡", Name: "因幡", Kokuryoku: 4, Region: 8, Neighbors: []string{"但馬", "伯耆", "美作", "播磨"}},
		"伯耆": {ID: "伯耆", Name: "伯耆", Kokuryoku: 12, Region: 8, Neighbors: []string{"因幡", "出雲", "備後", "美作", "備中"}},
		"出雲": {ID: "出雲", Name: "出雲", Kokuryoku: 6, Region: 8, Neighbors: []string{"伯耆", "備後", "石見"}},
		"石見": {ID: "石見", Name: "石見", Kokuryoku: 4, Region: 8, Neighbors: []string{"出雲", "安芸", "周防", "長門"}, GoldMine: true},

		// --- 地域 9 ---
		"備前": {ID: "備前", Name: "備前", Kokuryoku: 8, Region: 9, Neighbors: []string{"播磨", "美作", "備中"}},
		"備中": {ID: "備中", Name: "備中", Kokuryoku: 6, Region: 9, Neighbors: []string{"備前", "美作", "備後", "伯耆"}},
		"備後": {ID: "備後", Name: "備後", Kokuryoku: 6, Region: 9, Neighbors: []string{"備中", "安芸", "出雲", "伯耆", "石見"}},
		"安芸": {ID: "安芸", Name: "安芸", Kokuryoku: 6, Region: 9, Neighbors: []string{"備後", "周防", "石見"}},
		"美作": {ID: "美作", Name: "美作", Kokuryoku: 6, Region: 9, Neighbors: []string{"播磨", "備前", "備中", "伯耆", "因幡"}},
		"讃岐": {ID: "讃岐", Name: "讃岐", Kokuryoku: 4, Region: 9, Neighbors: []string{"阿波", "伊予"}},
		"伊予": {ID: "伊予", Name: "伊予", Kokuryoku: 10, Region: 9, Neighbors: []string{"讃岐", "土佐"}},

		// --- 地域 10 ---
		"紀伊": {ID: "紀伊", Name: "紀伊", Kokuryoku: 6, Region: 10, Neighbors: []string{"伊勢", "大和", "和泉", "河内"}},
		"阿波": {ID: "阿波", Name: "阿波", Kokuryoku: 6, Region: 10, Neighbors: []string{"讃岐", "土佐", "伊予"}},
		"土佐": {ID: "土佐", Name: "土佐", Kokuryoku: 6, Region: 10, Neighbors: []string{"阿波", "伊予"}},
		"日向": {ID: "日向", Name: "日向", Kokuryoku: 4, Region: 10, Neighbors: []string{"豊後", "肥後", "大隅"}},
		"薩摩": {ID: "薩摩", Name: "薩摩", Kokuryoku: 8, Region: 10, Neighbors: []string{"肥後", "大隅"}},
		"大隅": {ID: "大隅", Name: "大隅", Kokuryoku: 6, Region: 10, Neighbors: []string{"日向", "薩摩"}},

		// --- 地域 11 ---
		"肥後": {ID: "肥後", Name: "肥後", Kokuryoku: 12, Region: 11, Neighbors: []string{"筑後", "豊後", "日向", "薩摩"}},
		"筑前": {ID: "筑前", Name: "筑前", Kokuryoku: 12, Region: 11, Neighbors: []string{"豊前", "肥前", "筑後", "豊後"}, TradePort: true},
		"豊前": {ID: "豊前", Name: "豊前", Kokuryoku: 6, Region: 11, Neighbors: []string{"筑前", "長門", "豊後"}},
		"肥前": {ID: "肥前", Name: "肥前", Kokuryoku: 10, Region: 11, Neighbors: []string{"筑前", "筑後"}, TradePort: true},
		"筑後": {ID: "筑後", Name: "筑後", Kokuryoku: 8, Region: 11, Neighbors: []string{"筑前", "肥前", "豊後", "肥後"}},
		"豊後": {ID: "豊後", Name: "豊後", Kokuryoku: 12, Region: 11, Neighbors: []string{"豊前", "筑後", "肥後", "日向", "筑前"}, TradePort: true},
	}

	return provinces
}
