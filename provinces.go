package main

func InitializeProvinces() map[string]*Province {
	provinces := map[string]*Province{
		// --- 地域 1 (東北) ---
		"mutsu_kita":   {ID: "mutsu_kita", Name: "陸奥(北)", Kokuryoku: 6, Region: 1, Neighbors: []string{"mutsu_naka"}},
		"mutsu_naka":   {ID: "mutsu_naka", Name: "陸奥(中)", Kokuryoku: 12, Region: 1, Neighbors: []string{"mutsu_kita", "mutsu_minami"}},           // 出羽とは山脈で遮断
		"mutsu_minami": {ID: "mutsu_minami", Name: "陸奥(南)", Kokuryoku: 16, Region: 1, Neighbors: []string{"mutsu_naka", "kozuke", "shimotsuke"}}, // 常陸とは山脈で遮断
		"dewa":         {ID: "dewa", Name: "出羽", Kokuryoku: 12, Region: 1, Neighbors: []string{"kozuke"}},                                        // 陸奥、越後とは山脈で遮断

		// --- 地域 2 (坂東) ---
		"musashi": {ID: "musashi", Name: "武蔵", Kokuryoku: 16, Region: 2, Neighbors: []string{"sagami", "kozuke", "shimosa"}}, // 甲斐とは山脈
		"sagami":  {ID: "sagami", Name: "相模", Kokuryoku: 6, Region: 2, Neighbors: []string{"musashi", "suruga", "izu"}},
		"kazusa":  {ID: "kazusa", Name: "上総", Kokuryoku: 10, Region: 2, Neighbors: []string{"awa", "shimosa"}},
		"awa":     {ID: "awa", Name: "安房", Kokuryoku: 2, Region: 2, Neighbors: []string{"kazusa"}},
		"izu":     {ID: "izu", Name: "伊豆", Kokuryoku: 2, Region: 2, Neighbors: []string{"sagami", "suruga"}},

		// --- 地域 3 (関東) ---
		"shimotsuke": {ID: "shimotsuke", Name: "下野", Kokuryoku: 10, Region: 3, Neighbors: []string{"mutsu_minami", "kozuke", "hitachi", "shimosa"}},
		"kozuke":     {ID: "kozuke", Name: "上野", Kokuryoku: 12, Region: 3, Neighbors: []string{"mutsu_minami", "dewa", "shimotsuke", "musashi", "shinano_kita"}},
		"hitachi":    {ID: "hitachi", Name: "常陸", Kokuryoku: 14, Region: 3, Neighbors: []string{"shimotsuke", "shimosa"}},
		"shimosa":    {ID: "shimosa", Name: "下総", Kokuryoku: 10, Region: 3, Neighbors: []string{"hitachi", "shimotsuke", "musashi", "kazusa"}},

		// --- 地域 4 (北陸) ---
		"echigo":  {ID: "echigo", Name: "越後", Kokuryoku: 14, Region: 4, Neighbors: []string{"shinano_kita", "shinano_minami"}}, // 出羽、越中とは山脈
		"sado":    {ID: "sado", Name: "佐渡", Kokuryoku: 2, Region: 4, Neighbors: []string{}, GoldMine: true},                    // 島
		"etchu":   {ID: "etchu", Name: "越中", Kokuryoku: 12, Region: 4, Neighbors: []string{"noto", "kaga", "shinano_kita"}},    // 越後、飛騨とは山脈
		"noto":    {ID: "noto", Name: "能登", Kokuryoku: 6, Region: 4, Neighbors: []string{"etchu", "kaga"}},
		"kaga":    {ID: "kaga", Name: "加賀", Kokuryoku: 10, Region: 4, Neighbors: []string{"etchu", "noto", "echizen"}, Ikko: true},
		"echizen": {ID: "echizen", Name: "越前", Kokuryoku: 14, Region: 4, Neighbors: []string{"kaga", "wakasa"}, Ikko: true}, // 近江とは山脈

		// --- 地域 5 (中部) ---
		"shinano_kita":   {ID: "shinano_kita", Name: "信濃(北)", Kokuryoku: 8, Region: 5, Neighbors: []string{"kozuke", "echigo", "etchu", "shinano_minami"}}, // 飛騨とは山脈
		"shinano_minami": {ID: "shinano_minami", Name: "信濃(南)", Kokuryoku: 6, Region: 5, Neighbors: []string{"shinano_kita", "echigo", "kai", "mikawa", "owari", "mino"}},
		"hida":           {ID: "hida", Name: "飛騨", Kokuryoku: 2, Region: 5, Neighbors: []string{}, GoldMine: true},          // 四方を山脈に囲まれている？（地図上は山脈記号あり）
		"kai":            {ID: "kai", Name: "甲斐", Kokuryoku: 6, Region: 5, Neighbors: []string{"shinano_minami", "suruga"}}, // 武蔵とは山脈
		"suruga":         {ID: "suruga", Name: "駿河", Kokuryoku: 6, Region: 5, Neighbors: []string{"sagami", "izu", "kai", "totomi"}},
		"totomi":         {ID: "totomi", Name: "遠江", Kokuryoku: 6, Region: 5, Neighbors: []string{"suruga", "mikawa"}},
		"mikawa":         {ID: "mikawa", Name: "三河", Kokuryoku: 8, Region: 5, Neighbors: []string{"totomi", "shinano_minami", "owari"}},

		// --- 地域 6 (畿内・東海) ---
		"owari":      {ID: "owari", Name: "尾張", Kokuryoku: 14, Region: 6, Neighbors: []string{"mikawa", "shinano_minami", "mino", "ise"}},
		"mino":       {ID: "mino", Name: "美濃", Kokuryoku: 14, Region: 6, Neighbors: []string{"owari", "shinano_minami", "omi_kita", "ise"}},       // 飛騨、近江南とは山脈
		"omi_kita":   {ID: "omi_kita", Name: "近江(北)", Kokuryoku: 10, Region: 6, Neighbors: []string{"mino", "omi_minami", "yamashiro", "wakasa"}}, // 越前とは山脈
		"omi_minami": {ID: "omi_minami", Name: "近江(南)", Kokuryoku: 14, Region: 6, Neighbors: []string{"omi_kita", "yamashiro", "iga", "ise"}},     // 美濃とは山脈
		"ise":        {ID: "ise", Name: "伊勢", Kokuryoku: 18, Region: 6, Neighbors: []string{"owari", "mino", "omi_minami", "iga", "yamato", "kii"}},
		"iga":        {ID: "iga", Name: "伊賀", Kokuryoku: 2, Region: 6, Neighbors: []string{"ise", "omi_minami", "yamato"}},
		"yamato":     {ID: "yamato", Name: "大和", Kokuryoku: 12, Region: 6, Neighbors: []string{"ise", "iga", "kii", "kawachi"}},
		"kii":        {ID: "kii", Name: "紀伊", Kokuryoku: 6, Region: 6, Neighbors: []string{"ise", "yamato", "izumi"}},
		"yamashiro":  {ID: "yamashiro", Name: "山城", Kokuryoku: 14, Region: 6, Neighbors: []string{"omi_kita", "omi_minami", "settsu", "kawachi", "tamba"}},
		"settsu":     {ID: "settsu", Name: "摂津", Kokuryoku: 10, Region: 6, Neighbors: []string{"yamashiro", "izumi", "kawachi", "harima", "tamba"}, Honganji: true},
		"kawachi":    {ID: "kawachi", Name: "河内", Kokuryoku: 6, Region: 6, Neighbors: []string{"yamashiro", "yamato", "settsu", "izumi"}},
		"izumi":      {ID: "izumi", Name: "和泉", Kokuryoku: 2, Region: 6, Neighbors: []string{"settsu", "kawachi", "kii"}},

		// --- 地域 7 (山陽) ---
		"harima":  {ID: "harima", Name: "播磨", Kokuryoku: 12, Region: 7, Neighbors: []string{"settsu", "tamba", "bizen", "mimasaka"}},
		"bizen":   {ID: "bizen", Name: "備前", Kokuryoku: 8, Region: 7, Neighbors: []string{"harima", "mimasaka", "bit_chu"}},
		"bit_chu": {ID: "bit_chu", Name: "備中", Kokuryoku: 6, Region: 7, Neighbors: []string{"bizen", "mimasaka", "bingo"}},
		"bingo":   {ID: "bingo", Name: "備後", Kokuryoku: 6, Region: 7, Neighbors: []string{"bit_chu", "aki"}}, // 備中、美作、伯耆とは山脈
		"aki":     {ID: "aki", Name: "安芸", Kokuryoku: 6, Region: 7, Neighbors: []string{"bingo", "suo"}},
		"suo":     {ID: "suo", Name: "周防", Kokuryoku: 6, Region: 7, Neighbors: []string{"aki", "nagato"}},
		"nagato":  {ID: "nagato", Name: "長門", Kokuryoku: 4, Region: 7, Neighbors: []string{"suo"}},

		// --- 地域 8 (山陰) ---
		"wakasa": {ID: "wakasa", Name: "若狭", Kokuryoku: 2, Region: 8, Neighbors: []string{"echizen", "omi_kita", "tango"}},
		"tango":  {ID: "tango", Name: "丹後", Kokuryoku: 4, Region: 8, Neighbors: []string{"wakasa", "tajima", "tamba"}},
		"tajima": {ID: "tajima", Name: "但馬", Kokuryoku: 4, Region: 8, Neighbors: []string{"tango", "inaba", "tamba"}},
		"inaba":  {ID: "inaba", Name: "因幡", Kokuryoku: 4, Region: 8, Neighbors: []string{"tajima", "hoki"}},
		"hoki":   {ID: "hoki", Name: "伯耆", Kokuryoku: 12, Region: 8, Neighbors: []string{"inaba", "izumo"}, GoldMine: true}, // 備後とは山脈
		"izumo":  {ID: "izumo", Name: "出雲", Kokuryoku: 6, Region: 8, Neighbors: []string{"hoki", "iwami"}, GoldMine: true},
		"iwami":  {ID: "iwami", Name: "石見", Kokuryoku: 4, Region: 8, Neighbors: []string{"izumo"}, GoldMine: true},

		// 中間領域 (tamba, mimasaka)
		"tamba":    {ID: "tamba", Name: "丹波", Kokuryoku: 6, Region: 8, Neighbors: []string{"yamashiro", "settsu", "harima", "tango", "tajima"}},
		"mimasaka": {ID: "mimasaka", Name: "美作", Kokuryoku: 6, Region: 7, Neighbors: []string{"harima", "bizen", "bit_chu"}}, // 伯耆、備後とは山脈

		// --- 地域 9 (四国) ---
		"awa_shikoku": {ID: "awa_shikoku", Name: "阿波", Kokuryoku: 6, Region: 9, Neighbors: []string{"sanuki", "tosa"}},
		"sanuki":      {ID: "sanuki", Name: "讃岐", Kokuryoku: 4, Region: 9, Neighbors: []string{"awa_shikoku", "iyo"}},
		"tosa":        {ID: "tosa", Name: "土佐", Kokuryoku: 6, Region: 9, Neighbors: []string{"awa_shikoku", "iyo"}},
		"iyo":         {ID: "iyo", Name: "伊予", Kokuryoku: 10, Region: 9, Neighbors: []string{"sanuki", "tosa"}},

		// --- 地域 10 (南九州) ---
		"higo":    {ID: "higo", Name: "肥後", Kokuryoku: 12, Region: 10, Neighbors: []string{"chikugo", "bungo", "hyuga", "satsuma"}},
		"hyuga":   {ID: "hyuga", Name: "日向", Kokuryoku: 4, Region: 10, Neighbors: []string{"bungo", "higo", "osumi"}},
		"satsuma": {ID: "satsuma", Name: "薩摩", Kokuryoku: 8, Region: 10, Neighbors: []string{"higo", "osumi"}},
		"osumi":   {ID: "osumi", Name: "大隅", Kokuryoku: 6, Region: 10, Neighbors: []string{"hyuga", "satsuma"}},

		// --- 地域 11 (北九州) ---
		"chikuzen": {ID: "chikuzen", Name: "筑前", Kokuryoku: 12, Region: 11, Neighbors: []string{"buzen", "hizen", "chikugo"}},
		"buzen":    {ID: "buzen", Name: "豊前", Kokuryoku: 6, Region: 11, Neighbors: []string{"chikuzen", "chikugo", "bungo"}},
		"hizen":    {ID: "hizen", Name: "肥前", Kokuryoku: 10, Region: 11, Neighbors: []string{"chikuzen", "chikugo"}},
		"chikugo":  {ID: "chikugo", Name: "筑後", Kokuryoku: 8, Region: 11, Neighbors: []string{"chikuzen", "buzen", "hizen", "bungo", "higo"}},
		"bungo":    {ID: "bungo", Name: "豊後", Kokuryoku: 12, Region: 11, Neighbors: []string{"buzen", "chikugo", "higo", "hyuga"}},
	}

	return provinces
}
