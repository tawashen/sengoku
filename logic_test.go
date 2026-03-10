// package main

// import (
// 	"testing"
// )

// func TestRoleOrder(t *testing.T) {
// 	gs := &GameState{
// 		Players: []*Player{
// 			{ID: "1", Name: "Player 1"},
// 			{ID: "2", Name: "Player 2"},
// 			{ID: "3", Name: "Player 3"},
// 		},
// 	}
// 	m := model{gameState: gs}

// 	// 実行
// 	m = m.RoleOrder()

// 	// 検証: 全員に1〜Nのいずれかが割り当てられているか
// 	n := len(gs.Players)
// 	found := make(map[int]bool)
// 	for _, p := range gs.Players {
// 		if p.Order < 1 || p.Order > n {
// 			t.Errorf("Player %s has invalid order: %d", p.ID, p.Order)
// 		}
// 		if found[p.Order] {
// 			t.Errorf("Duplicate order found: %d", p.Order)
// 		}
// 		found[p.Order] = true
// 	}

// 	// 検証: 全ての数字(1〜N)が使われているか
// 	if len(found) != n {
// 		t.Errorf("Expected %d unique orders, got %d", n, len(found))
// 	}
// }
