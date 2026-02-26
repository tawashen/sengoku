package main

// GamePhase represents the current stage of the turn
type GamePhase int

const (
	TaxPhase GamePhase = iota
	ArmyPhase
	DiplomacyPhase
	MarchPhase
	MaintenancePhase
)

func (p GamePhase) String() string {
	return []string{"徴税", "軍備", "外交・調略", "行軍・合戦", "俸禄・維持"}[p]
}

// Province represents a country on the map
type Province struct {
	ID          string
	Name        string
	Kokudaka    int      // 石高
	OwnerID     string   // 所属プレイヤーID
	Castles     int      // 城の数
	Soldiers    int      // 兵士数
	IsRestless  bool     // 不穏状態
	HasUprising bool     // 一揆発生中
	Neighbors   []string // 隣接する国のID (つながり)
}

// General represents a Sengoku Daimyo or vassal
type General struct {
	ID         string
	Name       string
	Combat     int    // 戦闘能力
	Politics   int    // 内政能力
	Ambition   int    // 野心
	Loyalty    int    // 忠誠度
	Stipend    int    // 俸禄
	ProvinceID string // 所在国ID
	OwnerID    string // 所属プレイヤーID
}

// Player represents a clan/daimyo controller
type Player struct {
	ID   string
	Name string
	Gold int
	Clan string
	IsAI bool
}

// GameState holds the entire game world data
type GameState struct {
	Year      int
	Phase     GamePhase
	Provinces map[string]*Province
	Generals  map[string]*General
	Players   map[string]*Player
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
