package main

func InitializeProvinces() map[string]*Province {
	provinces := map[string]*Province{
		"owari": {
			ID:        "owari",
			Name:      "尾張",
			Kokudaka:  10,
			Neighbors: []string{"mikawa", "mino", "ise"},
		},
		"mikawa": {
			ID:        "mikawa",
			Name:      "三河",
			Kokudaka:  5,
			Neighbors: []string{"owari", "totomi", "shinano"},
		},
		"mino": {
			ID:        "mino",
			Name:      "美濃",
			Kokudaka:  8,
			Neighbors: []string{"owari", "omi", "echizen", "shinano"},
		},
	}
	return provinces
}
