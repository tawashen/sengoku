package main

// Province represents a country on the map
type Province struct {
	ID       string
	Name     string
	Kokudaka int    // 石高
	OwnerID  string // 所属プレイヤーID
	Complete bool   //完全支配かどうか
	Castles  []Castle
	Soldiers int  // 兵士数
	Restless bool // 不穏状態
	//HasUprising bool     // 一揆発生
	Starving  bool //飢饉発生
	Christian bool //キリシタン発生
	TradePort bool
	GoldMine  bool
	Neighbors []string // 隣接する国のID (つながり)
}

// General represents a Sengoku Daimyo or vassal
type General struct {
	ID           string //モブ武将もいる
	Name         string
	Combat       int    // 戦闘能力
	PlusCombat   int    //戦闘能力補正値
	Politics     int    // 内政能力
	Prestige     int    //威信
	PlusPrestige int    //威信補正
	Loyalty      int    // 忠誠度
	PlusLoyalty  int    //忠誠度補正
	Stipend      int    // 俸禄
	ProvinceID   string // 所在国ID
	OwnerID      string // 所属プレイヤーID
}

type Castle struct {
	Ruler            string //Player ID　もしくは中立
	Power            int    //1につき1000人
	IkkouUprising    bool   //一揆発生
	DoUprising       bool   //土一揆
	ProvinceUprising bool   //国一揆
	Isolated         bool   //孤立中
	Surrounded       bool   //包囲中
}

// Player represents a clan/daimyo controller
type Player struct {
	ID        string
	Name      string //大名の名前
	Gold      int
	Clan      string
	IsAI      bool
	Generals  []string //大名の下の将校
	Provinces []string //大名の所持する国
	Power     int      //国力の合計
	//Order     int      // 順番
}

type Card struct {
	Name        string //キーはこれを使う
	Description string //説明内容
	Secret      bool   //秘密
	Event       bool   //事件
	Dice        *Dice  //これをキーにして辞書内に入れたメソッドを呼び出す？
}

type Dice struct {
	Result [6]any
}

// GameState holds the entire game world data
type GameState struct {
	Year      int
	Phase     string
	Provinces map[string]*Province
	Generals  map[string]*General
	Players   []*Player
	Order     []int //PlayerのIndex用
	Cards     []Card
}

// Helper: Check if two provinces are neighbors
func (gs *GameState) AreNeighbors(id1, id2 string) bool {
	p1, ok := gs.Provinces[id1]
	if !ok {
		return false
	}
	for _, neighborID := range p1.Neighbors {
		if neighborID == id2 {
			return true
		}
	}
	return false
}
