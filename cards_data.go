package main

func InitializeCards() []Card {
	return []Card{
		// --- card1.jpg ---

		{
			Name:        "国内の不穏",
			Description: "このカードをひいた大名はこの回に徴税を行う支配地域は自動的に不穏となり、それぞれについて一揆チェックを行う。徴税を行わない支配地域は、不穏とはならない。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true, //徴税時に使用
		},
		{
			Name:        "一向一揆",
			Description: "サイコロを振って、発生する地域を決定する。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{"越前", "加賀", "三河", "越中", "伊勢", "本願寺"},
			},
		},
		{
			Name:        "廃鉱",
			Description: "支配地域内の金山のうち一つ（任意）からは、以後ゲームを通じて、金を受けとることはできない。（金山マーカーを取り去る。）",
			Secret:      false,
			Event:       true,
		},
		{
			Name:        "裏切り",
			Description: "現在の俸禄が最も少ない家臣が裏切る。",
			Secret:      false,
			Event:       true,
		},
		{
			Name:        "大名死亡",
			Description: "新当主（世継ぎ）をたてた後、全家臣について裏切りチェックを行う。",
			Secret:      false,
			Event:       true,
		},
		{
			Name:        "内応の露顕",
			Description: "「城方の内応」のカードを無効にできる。攻撃側は強襲を行わなければならない。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "城方の内応Auto",
			Description: "攻城戦の際にこのカードを出すと、目標の城塞を取り除ける。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "城方の内応Dice",
			Description: "サイコロを一つ振り、下の表にしたがって、相手の城のレベル数を低下させることができる。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{"5戦力", "10戦力", "10戦力", "15戦力", "15戦力", "20戦力"},
			},
		},
		{
			Name:        "城方の内応Dice",
			Description: "サイコロを一つ振り、下の表にしたがって、相手の城のレベル数を低下させることができる。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{"5戦力", "10戦力", "15戦力", "20戦力", "25戦力", "30戦力"},
			},
		},
		{
			Name:        "詭計の成功",
			Description: "野戦の際、相手の戦力を二等分し、そのそれぞれと全戦力で戦える。それぞれの最初の戦闘ラウンドでは、サイコロの目が3有利になる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "伏兵",
			Description: "防御側のみ野戦の最初の戦闘ラウンドで、戦闘解決のサイコロの目が5有利になる。攻撃側の「奇襲」「詭計の成功」カードは無効となる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "奇襲",
			Description: "相手に野戦を強制できる。最初の戦闘ラウンドで戦闘解決のサイコロの目が5有利になる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "奇襲",
			Description: "相手に野戦を強制できる。最初の戦闘ラウンドで戦闘解決のサイコロの目が4有利になる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "飢饉",
			Description: "陸奥、出羽全域に飢饉が発生する。混乱状態となり、徴税できない。隣接する地域毎にサイコロを振り、5か6が出るとその地域にも飢饉が発生する。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "豊作",
			Description: "サイコロを振り、出た目に対応する地域が豊作となる。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true,
			Dice: &Dice{
				Result: [6]any{5, 6, 7, []int{8, 9}, 10, 11},
			},
		},

		// --- card2.jpg ---

		{
			Name:        "影武者",
			Description: "戦闘、刺客による大名・家臣の死亡をたわがわりできる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "反対勢力の増大",
			Description: "このカードをひいた大名のすべての支配国で、中立勢力の発生チェックを行う。サイコロを振り 完全支配国 6で発生 部分支配国 5,6で発生。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "凶作",
			Description: "サイコロを振り、出た目に対応する地域が凶作となる。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true,
			Dice: &Dice{
				Result: [6]any{1, 1, 2, 3, 4, 5},
			},
		},
		{
			Name:        "凶作",
			Description: "サイコロを振り、出た目に対応する地域が凶作となる。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true,
			Dice: &Dice{
				Result: [6]any{6, 7, 8, 9, 10, 11},
			},
		},
		{
			Name:        "南蛮人の渡来",
			Description: "サイコロを振り、来航した国に「南蛮貿易港マーカー」を置く。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{"和泉", "筑前", "肥前", "豊後", "肥後", "薩摩"},
			},
		},
		{
			Name:        "築城技術の進歩",
			Description: "支配地域内にある城のうち1ヶ所を、1レベル（5戦力分）大きくすることができる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "九州での反対勢力の増大",
			Description: "九州各地で、中立勢力の発生チェックを行なう。サイコロを振り 完全支配国 6で発生 部分支配国 5,6で発生 非支配国 4,5,6で発生。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "土一揆との和睦",
			Description: "自分の領国内での土一揆の発生を無効にできる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "威信の上昇",
			Description: "大名の威信が1上昇するが、朝廷への献金として 金20 を支払わなくてはならない。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "隣国での混乱",
			Description: "自分の勢力が存在する国に隣接する国に、中立勢力が発生する。（一国のみ）サイコロを一つ振り、大名の威信以下の目が出れば、発生しその中立勢力のうち5戦力を自分の勢力にできる。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "キリシタンの流行",
			Description: "南蛮貿易マーカーのある国すべてに「キリシタンマーカー」を置く。隣接する国毎にサイコロを振り、5,6が出たら、その国にも「キリシタンマーカー」を置く。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "他の勢力の調整",
			Description: "自分の支配地域内または支配地域の隣国にある他の勢力のうち、5戦力を自分の勢力にできる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "金山発見",
			Description: "金山のマークがある国を支配している大名は、このカードを出したターン以降、毎年指定された額の金を得られる。（金山マーカーを置く。）",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "豊作",
			Description: "サイコロを振り、出た目に対応する地域が豊作となる。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Tax:         true,
			Dice: &Dice{
				Result: [6]any{[]int{10, 11}, 1, []int{2, 3}, 4, 5, []int{6, 7}},
			},
		},
		{
			Name:        "国一揆との和睦",
			Description: "自分の領国内での国一揆の発生を無効にできる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "一向宗との和睦",
			Description: "自分の領国内での一向一揆の発生を無効にできる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "刺客",
			Description: "他家の武将のうち1人を指名し、暗殺のチェックを行う。依頼額：金10。サイコロを振り、大名暗殺：1, 2で成功、家臣暗殺：1, 2, 3で成功。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "大義",
			Description: "将軍から御内書が届く。山城を支配している大名は10勝利得点をボーナスとして得られる。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "関東での反対勢力の増大",
			Description: "関東全域で、中立勢力の発生チェックを行う。サイコロを振り 完全支配国 6で発生 部分支配国 5,6で発生 非支配国 4,5,6で発生。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "陸奥・出羽での反対勢力の増大",
			Description: "陸奥・出羽で、中立勢力の発生チェックを行う。サイコロを振り 完全支配国 6で発生 部分支配国 5,6で発生 非支配国 4,5,6で発生。",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "敵方家臣の調略",
			Description: "他の大名の家臣に対して、調略を行う。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "豊作",
			Description: "サイコロを振り、出た目に対応する地域が豊作となる。(このカードは、徴税フェイズに表にする)",
			Secret:      false,
			Event:       true,
			Dice: &Dice{
				Result: [6]any{8, []int{9, 10}, []int{11, 1}, 2, 3, 4},
			},
		},
		{
			Name:        "下剋上",
			Description: "家臣のうち最も優秀なものが大名にとってかわる。他の家臣それぞれについて裏切りチェックを行い、残った家臣に対し、必ず俸禄の倍増を行う。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "富の浪費",
			Description: "雅びな暮らしに溺れ、所持している金の半分（端数切り捨て）を失う。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "不穏工作",
			Description: "自分の武将がいる国またはその隣国のうち一つを不穏にできる。サイコロを一つ振り、工作を行う武将の威信以下の数が出れば成功する。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "家臣死亡",
			Description: "サイコロを一つ振り、どの家臣が死亡するかを決定する。",
			Secret:      true,
			Event:       false,
			Dice: &Dice{
				Result: [6]any{1, 2, 3, 4, 5, 6},
			},
		},
		{
			Name:        "密報",
			Description: "他のプレイヤーが持っているカードを1枚見ることができる。それが戦闘カードなら、捨てなければならない。",
			Secret:      true,
			Event:       false,
		},
		{
			Name:        "賢臣の諫言",
			Description: "自分の家臣や大名に対して凶となるカードをひいた場合、それを無効にできる。",
			Secret:      true,
			Event:       false,
		},
	}
}
