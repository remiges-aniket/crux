// /*
// This file contains the functions that represent BRE tests for doMatch(). These functions are called
// inside TestDoMatch() in do_match_test.go.

// Some of the definitions of rulesets below deliberately use a lot of whitespace to keep the code consistent
// and to make it easier to understand, add to, and edit these tests
// */

package crux

// const (
// 	// The "main" ruleset that may contain "thenCall"s/"elseCall"s to other rulesets
// 	mainRS = "main"

// 	inventoryItemClass = "inventoryitem"
// 	transactionClass   = "transaction"
// 	purchaseClass      = "purchase"
// 	orderClass         = "order"
// )

// /*
// Realm:   1,

// 	App:     "Test5",
// 	Slice:   6,
// 	Class:   "inventoryitem",
// */
// var sampleEntity = Entity{
// 	Realm: "1",
// 	App:   "Test6",
// 	Slice: "6",
// 	Class: inventoryItemClass,
// 	Attrs: map[string]string{
// 		"cat":        "textbook",
// 		"fullname":   "Advanced Physics",
// 		"ageinstock": "5",
// 		"mrp":        "50.80",
// 		"received":   "2018-06-01T15:04:05Z",
// 		"bulkorder":  trueStr,
// 	},
// }

// func testBasic(tests *[]doMatchTest) {

// 	rule1 := Rule_t{

// 		RulePatterns: []RulePatternBlock_t{{"cat", opEQ, "textbook"}},
// 		RuleActions: RuleActionBlock_t{
// 			Task:       []string{"yearendsale", "summersale"},
// 			Properties: map[string]string{"cashback": "10", "discount": "9"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"yearendsale", "summersale"},
// 		Properties: map[string]string{"cashback": "10", "discount": "9"},
// 	}

// 	*tests = append(*tests, doMatchTest{
// 		"basic test", sampleEntity, &rs, ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testExit(tests *[]doMatchTest) {

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,

// 		NCalled: 0,
// 	}

// 	rs.Rules = []Rule_t{
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"cat", opEQ, "refbook"}},
// 			RuleActions: RuleActionBlock_t{
// 				Task:       []string{"springsale"},
// 				Properties: map[string]string{"cashback": "15"},
// 			}, // no match
// 		},
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"ageinstock", opLT, 7}, {"cat", opEQ, "textbook"}},
// 			RuleActions: RuleActionBlock_t{
// 				Task:       []string{"yearendsale", "summersale"},
// 				Properties: map[string]string{"discount": "10", "freegift": "mug"},
// 			}, // match
// 		},
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"summersale", opEQ, true}},
// 			RuleActions: RuleActionBlock_t{
// 				Task:       []string{"wintersale"},
// 				Properties: map[string]string{"discount": "15"},
// 				DoExit:     true,
// 			}, // match then exit
// 		},
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"ageinstock", opLT, 7}},
// 			RuleActions: RuleActionBlock_t{
// 				Task: []string{"autumnsale"},
// 			}, // ignored
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"yearendsale", "summersale", "wintersale"},
// 		Properties: map[string]string{"discount": "15", "freegift": "mug"},
// 	}

// 	*tests = append(*tests, doMatchTest{"exit", sampleEntity, &rs, ActionSet{
// 		Tasks:      []string{},
// 		Properties: make(map[string]string),
// 	}, want})
// }

// func testReturn(tests *[]doMatchTest) {

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,

// 		NCalled: 0,
// 	}

// 	rs.Rules = []Rule_t{
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"ageinstock", opLT, 7}, {"cat", opEQ, "textbook"}},
// 			RuleActions: RuleActionBlock_t{
// 				Task:       []string{"yearendsale", "summersale"},
// 				Properties: map[string]string{"discount": "10", "freegift": "mug"},
// 			}, // match then exit

// 		},
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"summersale", opEQ, true}},
// 			RuleActions:  RuleActionBlock_t{Task: []string{"springsale"}, Properties: map[string]string{"discount": "15"}, DoReturn: true}, // match then return
// 		},
// 		{
// 			RulePatterns: []RulePatternBlock_t{{"ageinstock", opLT, 7}},
// 			RuleActions:  RuleActionBlock_t{Task: []string{"autumnsale"}}, // ignored
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"yearendsale", "summersale", "springsale"},
// 		Properties: map[string]string{"discount": "15", "freegift": "mug"},
// 	}

// 	*tests = append(*tests, doMatchTest{"return", sampleEntity, &rs, ActionSet{
// 		Tasks:      []string{},
// 		Properties: make(map[string]string),
// 	}, want})
// }

// func testTransactions(tests *[]doMatchTest) {

// 	ruleSetsTests = []*Ruleset_t{}
// 	ruleSetsTests = append(ruleSetsTests, &Ruleset_t{})

// 	rs := &Schema_t{
// 		Class: transactionClass,
// 		PatternSchema: []PatternSchema_t{

// 			{Attr: "productname", ValType: typeStr},
// 			{Attr: "price", ValType: typeInt},
// 			{Attr: "inwintersale", ValType: typeBool},
// 			{Attr: "paymenttype", ValType: typeEnum},
// 			{Attr: "ismember", ValType: typeBool},
// 		},
// 		ActionSchema: ActionSchema_t{},
// 	}
// 	ruleSchemasTest = append(ruleSchemasTest, rs)
// 	setupRuleSetsForTransaction()

// 	// Each test below involves calling doMatch() with a different entity
// 	testWinterDiscJacket60(tests)
// 	testWinterDiscJacket40(tests)
// 	testWinterDiscKettle110Cash(tests)
// 	testWinterDiscKettle110Card(tests)
// 	testMemberDiscLamp60(tests)
// 	testMemberDiscKettle60Card(tests)
// 	testMemberDiscKettle60Cash(tests)
// 	testMemberDiscKettle110Card(tests)
// 	testMemberDiscKettle110Cash(tests)
// 	testNonMemberDiscLamp30(tests)
// 	testNonMemberDiscKettle70(tests)
// 	testNonMemberDiscKettle110Cash(tests)
// 	testNonMemberDiscKettle110Card(tests)
// }

// func setupRuleSetsForTransaction() {
// 	setupRuleSetMainForTransaction()
// 	setupRuleSetWinterDisc()
// 	setupRuleSetRegularDisc()
// 	setupRuleSetMemberDisc()
// 	setupRuleSetNonMemberDisc()
// }

// func setupRuleSetMainForTransaction() *Ruleset_t {
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"inwintersale", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			ThenCall: "winterdisc",
// 			ElseCall: "regulardisc",
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"paymenttype", opEQ, "cash"},
// 			{"price", opGT, 10},
// 		},
// 		RuleActions: RuleActionBlock_t{Task: []string{"freepen"}},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"paymenttype", opEQ, "card"},
// 			{"price", opGT, 10},
// 		},
// 		RuleActions: RuleActionBlock_t{Task: []string{"freemug"}},
// 	}
// 	rule4 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"freehat", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{Task: []string{"freebag"}},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3, rule4},
// 		NCalled: 0,
// 	}
// 	return &rs

// }

// func setupRuleSetWinterDisc() {

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "jacket"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freehat", "freemug", "freebag"},
// 			Properties: map[string]string{"discount": "50"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opLT, 100},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "40", "pointsmult": "2"},
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opGE, 100},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "45", "pointsmult": "3"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3},
// 		NCalled: 0,
// 	}

// 	ruleSetsTests = append(ruleSetsTests, &rs)

// }

// func setupRuleSetRegularDisc() {

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"ismember", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			ThenCall: "memberdisc",
// 			ElseCall: "nonmemberdisc",
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	ruleSetsTests = append(ruleSetsTests, &rs)

// }

// func setupRuleSetMemberDisc() {

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "lamp"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "35", "pointsmult": "2"},
// 			DoExit:     true,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opLT, 100},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "20"},
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opGE, 100},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "25"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3},
// 		NCalled: 0,
// 	}

// 	ruleSetsTests = append(ruleSetsTests, &rs)

// }

// func setupRuleSetNonMemberDisc() {

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opLT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "5"},
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opGE, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "10"},
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opGE, 100},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "15"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3},
// 		NCalled: 0,
// 	}

// 	ruleSetsTests = append(ruleSetsTests, &rs)

// }

// func testWinterDiscJacket60(tests *[]doMatchTest) {

// 	ruleSetsTests = []*Ruleset_t{}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "jacket"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freehat", "freemug", "freebag"},
// 			Properties: map[string]string{"discount": "50"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	//ruleSetsTests = append(ruleSetsTests, &rs)

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "jacket",
// 			"price":        "60",
// 			"inwintersale": trueStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}
// 	want := ActionSet{
// 		Tasks:      []string{"freehat", "freemug", "freebag"},
// 		Properties: map[string]string{"discount": "50"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"winterdisc jacket 60",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testWinterDiscJacket40(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "jacket",
// 			"price":        "60",
// 			"inwintersale": trueStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "jacket"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freemug"},
// 			Properties: map[string]string{"discount": "40", "pointsmult": "2"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug"},
// 		Properties: map[string]string{"discount": "40", "pointsmult": "2"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"winterdisc jacket 40",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testWinterDiscKettle110Cash(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": trueStr,
// 			"paymenttype":  "cash",
// 			"ismember":     trueStr,
// 		},
// 	}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "45", "pointsmult": "3"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "45", "pointsmult": "3"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"winterdisc kettle 110 cash",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testWinterDiscKettle110Card(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": trueStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freemug"},
// 			Properties: map[string]string{"discount": "45", "pointsmult": "3"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug"},
// 		Properties: map[string]string{"discount": "45", "pointsmult": "3"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"winterdisc kettle 110 card",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testMemberDiscLamp60(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "lamp",
// 			"price":        "60",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "lamp"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{},
// 			Properties: map[string]string{"discount": "35", "pointsmult": "2"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"discount": "35", "pointsmult": "2"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"memberdisc lamp 60",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testMemberDiscKettle60Card(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "60",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freemug"},
// 			Properties: map[string]string{"discount": "35"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)

// 	want := ActionSet{
// 		Tasks:      []string{"freemug"},
// 		Properties: map[string]string{"discount": "35"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"memberdisc kettle 60 card",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testMemberDiscKettle60Cash(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "60",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "cash",
// 			"ismember":     trueStr,
// 		},
// 	}

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "20"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   inventoryItemClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "20"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"memberdisc kettle 60 cash",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testMemberDiscKettle110Card(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "card",
// 			"ismember":     trueStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freemug"},
// 			Properties: map[string]string{"discount": "25"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug"},
// 		Properties: map[string]string{"discount": "25"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"memberdisc kettle 110 card",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testMemberDiscKettle110Cash(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "cash",
// 			"ismember":     trueStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 50},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "25"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "25"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"memberdisc kettle 110 cash",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testNonMemberDiscLamp30(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "lamp",
// 			"price":        "30",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "cash",
// 			"ismember":     falseStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "lamp"},
// 			{"price", opGT, 25},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "5"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)
// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "5"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"nonmemberdisc lamp 30",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testNonMemberDiscKettle70(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "70",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "cash",
// 			"ismember":     falseStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 25},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "10"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)

// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "10"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"nonmemberdisc kettle 70",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testNonMemberDiscKettle110Cash(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "cash",
// 			"ismember":     falseStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 25},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freepen"},
// 			Properties: map[string]string{"discount": "15"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)
// 	want := ActionSet{
// 		Tasks:      []string{"freepen"},
// 		Properties: map[string]string{"discount": "15"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"nonmemberdisc kettle 110 cash",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testNonMemberDiscKettle110Card(tests *[]doMatchTest) {
// 	ruleSetsTests = []*Ruleset_t{}
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test7",
// 		Slice: "7",
// 		Class: transactionClass,
// 		Attrs: map[string]string{
// 			"productname":  "kettle",
// 			"price":        "110",
// 			"inwintersale": falseStr,
// 			"paymenttype":  "card",
// 			"ismember":     falseStr,
// 		},
// 	}
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"productname", opEQ, "kettle"},
// 			{"price", opGT, 25},
// 		},
// 		RuleActions: RuleActionBlock_t{

// 			Task:       []string{"freemug"},
// 			Properties: map[string]string{"discount": "15"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   transactionClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)
// 	want := ActionSet{
// 		Tasks:      []string{"freemug"},
// 		Properties: map[string]string{"discount": "15"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"nonmemberdisc kettle 110 card",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testPurchases(tests *[]doMatchTest) {
// 	//setupPurchaseRuleSchema()
// 	setupRuleSetForPurchases()

// 	// Each test below involves calling doMatch() with a different entity
// 	testJacket35(tests)
// 	testJacket55ForMember(tests)
// 	testJacket55ForNonMember(tests)
// 	testJacket75ForMember(tests)
// 	testJacket75ForNonMember(tests)
// 	testLamp35(tests)
// 	testLamp55(tests)
// 	testLamp75ForMember(tests)
// 	testLamp75ForNonMember(tests)
// 	testKettle35(tests)
// 	testKettle55(tests)
// 	testKettle75ForMember(tests)
// 	testKettle75ForNonMember(tests)
// 	testOven35(tests)
// 	testOven55(tests)
// }

// func setupPurchaseRuleSchema() {

// 	ruleSchemasTest = append(ruleSchemasTest, &Schema_t{
// 		Class: purchaseClass,
// 		PatternSchema: []PatternSchema_t{

// 			{Attr: "product", ValType: typeStr},
// 			{Attr: "price", ValType: typeFloat},
// 			{Attr: "ismember", ValType: typeBool},
// 		},
// 		ActionSchema: ActionSchema_t{
// 			Tasks: []string{"freepen", "freebottle", "freepencil", "freemug", "freejar", "freeplant",
// 				"freebag", "freenotebook"},
// 			Properties: []string{"discount", "pointsmult"},
// 		},
// 	})

// }

// func testJacket35(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "jacket",
// 			"price":    "35",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen", "freebottle", "freepencil"},
// 		Properties: map[string]string{"discount": "5"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"jacket price 35",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testJacket55ForMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "jacket",
// 			"price":    "55",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen", "freebottle", "freepencil", "freenotebook"},
// 		Properties: map[string]string{"discount": "10"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"jacket price 55 for member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testJacket55ForNonMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "jacket",
// 			"price":    "55",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen", "freebottle", "freepencil", "freenotebook"},
// 		Properties: map[string]string{"discount": "10"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"jacket price 55 for non-member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testJacket75ForMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "jacket",
// 			"price":    "75",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen", "freebottle", "freepencil", "freenotebook"},
// 		Properties: map[string]string{"discount": "15", "pointsmult": "2"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"jacket price 75 for member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testJacket75ForNonMember(tests *[]doMatchTest) {
// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "jacket",
// 			"price":    "75",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freepen", "freebottle", "freepencil", "freenotebook"},
// 		Properties: map[string]string{"discount": "10"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"jacket price 75 for non-member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testLamp35(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "lamp",
// 			"price":    "35",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug", "freejar", "freeplant", "freebag"},
// 		Properties: map[string]string{"discount": "20"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"lamp price 35",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testLamp55(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "lamp",
// 			"price":    "55",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug", "freejar", "freeplant", "freebag", "freenotebook"},
// 		Properties: map[string]string{"discount": "25"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"lamp price 55",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testLamp75ForMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "lamp",
// 			"price":    "75",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug", "freejar", "freeplant"},
// 		Properties: map[string]string{"discount": "30", "pointsmult": "3"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"lamp price 75 for member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testLamp75ForNonMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "lamp",
// 			"price":    "75",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freemug", "freejar", "freeplant", "freebag", "freenotebook"},
// 		Properties: map[string]string{"discount": "25"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"lamp price 75 for non-member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testKettle35(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "kettle",
// 			"price":    "35",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"discount": "35"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"kettle price 35",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testKettle55(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "kettle",
// 			"price":    "55",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freenotebook"},
// 		Properties: map[string]string{"discount": "40"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"kettle price 55",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testKettle75ForMember(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "kettle",
// 			"price":    "75",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"discount": "45", "pointsmult": "4"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"kettle price 75 for member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testKettle75ForNonMember(tests *[]doMatchTest) {
// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "kettle",
// 			"price":    "75",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{"freenotebook"},
// 		Properties: map[string]string{"discount": "40"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"kettle price 75 for non-member",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testOven35(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "oven",
// 			"price":    "35",
// 			"ismember": falseStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks:      []string{},
// 		Properties: make(map[string]string),
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"oven price 35",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testOven55(tests *[]doMatchTest) {

// 	rc := setupRuleSetForPurchases()
// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test8",
// 		Slice: "8",
// 		Class: purchaseClass,
// 		Attrs: map[string]string{
// 			"product":  "oven",
// 			"price":    "55",
// 			"ismember": trueStr,
// 		},
// 	}

// 	want := ActionSet{
// 		Tasks: []string{"freenotebook"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"oven price 55",
// 		entity,
// 		rc,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func setupRuleSetForPurchases() *Ruleset_t {
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "jacket"},
// 			{"price", opGT, 30.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Task:       []string{"freepen", "freebottle", "freepencil"},
// 			Properties: map[string]string{"discount": "5"},
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "jacket"},
// 			{"price", opGT, 50.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "10"},
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "jacket"},
// 			{"price", opGT, 70.0},
// 			{"ismember", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "15", "pointsmult": "2"},
// 		},
// 	}
// 	rule4 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "lamp"},
// 			{"price", opGT, 30.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Task:       []string{"freemug", "freejar", "freeplant"},
// 			Properties: map[string]string{"discount": "20"},
// 		},
// 	}
// 	rule5 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "lamp"},
// 			{"price", opGT, 50.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "25"},
// 		},
// 	}
// 	rule6 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "lamp"},
// 			{"price", opGT, 70.0},
// 			{"ismember", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "30", "pointsmult": "3"},
// 			DoExit:     true,
// 		},
// 	}
// 	rule7 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "kettle"},
// 			{"price", opGT, 30.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "35"},
// 		},
// 	}
// 	rule8 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "kettle"},
// 			{"price", opGT, 50.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "40"},
// 		},
// 	}
// 	rule9 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"product", opEQ, "kettle"},
// 			{"price", opGT, 70.0},
// 			{"ismember", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"discount": "45", "pointsmult": "4"},
// 			DoReturn:   true,
// 		},
// 	}
// 	rule10 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"freemug", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Task: []string{"freebag"},
// 		},
// 	}
// 	rule11 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"price", opGT, 50.0},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Task: []string{"freenotebook"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   purchaseClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3, rule4, rule5, rule6, rule7, rule8, rule9, rule10, rule11},
// 		NCalled: 0,
// 	}

// 	return &rs

// }

// func testOrders(tests *[]doMatchTest) {

// 	setupRuleSetMainForOrder()
// 	setupRuleSetPurchaseOrSIPForOrder()
// 	setupRuleSetOtherOrderTypesForOrder()

// 	// Each test below involves calling doMatch() with a different entity
// 	testSIPOrder(tests)
// 	testSwitchDematOrder(tests)
// 	testSwitchDematExtHours(tests)
// 	testRedemptionDematExtHours(tests)
// 	testPurchaseOvernightOrder(tests)
// 	testSIPLiquidOrder(tests)
// 	testSwitchPhysicalOrder(tests)
// }

// func setupRuleSetMainForOrder() *Ruleset_t {
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"ordertype", opEQ, "purchase"},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			ThenCall: "purchaseorsip",
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"ordertype", opEQ, "sip"},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			ThenCall: "purchaseorsip",
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"ordertype", opNE, "purchase"},
// 			{"ordertype", opNE, "sip"},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500"},
// 			ThenCall:   "otherordertypes",
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3},
// 		NCalled: 0,
// 	}
// 	return &rs
// }

// func setupRuleSetPurchaseOrSIPForOrder() *Ruleset_t {
// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"liquidscheme", opEQ, false},
// 			{"overnightscheme", opEQ, false},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1430",
// 				"fundscutoff": "1430"},
// 			DoReturn: true,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1330", "Novaordercutoff": "1300",
// 				"fundscutoff": "1230"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2},
// 		NCalled: 0,
// 	}

// 	return &rs
// }

// func setupRuleSetOtherOrderTypesForOrder() {

// 	rule1 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"mode", opEQ, "physical"},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Task: []string{"unitstoamc", "unitstorta"},
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"mode", opEQ, "demat"},
// 			{"extendedhours", opEQ, false},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"unitscutoff": "1630"},
// 		},
// 	}
// 	rule3 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{
// 			{"mode", opEQ, "demat"},
// 			{"extendedhours", opEQ, true},
// 		},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"unitscutoff": "1730"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule1, rule2, rule3},
// 		NCalled: 0,
// 	}

// 	ruleSetsTests = append(ruleSetsTests, &rs)
// }

// func testSIPOrder(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "sip",
// 			"mode":            "demat",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": falseStr,
// 			"extendedhours":   falseStr,
// 		},
// 	}

// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1430",
// 				"fundscutoff": "1430"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1430",
// 			"fundscutoff": "1430"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"sip order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testSwitchDematOrder(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "switch",
// 			"mode":            "demat",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": falseStr,
// 			"extendedhours":   falseStr,
// 		},
// 	}

// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 				"unitscutoff": "1630"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 			"unitscutoff": "1630"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"switch demat order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testSwitchDematExtHours(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "switch",
// 			"mode":            "demat",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": falseStr,
// 			"extendedhours":   trueStr,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 				"unitscutoff": "1730"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 			"unitscutoff": "1730"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"switch demat ext-hours order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testRedemptionDematExtHours(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "redemption",
// 			"mode":            "demat",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": falseStr,
// 			"extendedhours":   trueStr,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 				"unitscutoff": "1730"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1500", "Novaordercutoff": "1500",
// 			"unitscutoff": "1730"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"redemption demat ext-hours order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testPurchaseOvernightOrder(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "purchase",
// 			"mode":            "physical",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": trueStr,
// 			"extendedhours":   falseStr,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1330", "Novaordercutoff": "1300",
// 				"fundscutoff": "1230"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}

// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1330", "Novaordercutoff": "1300",
// 			"fundscutoff": "1230"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"purchase overnight order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }

// func testSIPLiquidOrder(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,
// 		Attrs: map[string]string{
// 			"ordertype":       "sip",
// 			"mode":            "physical",
// 			"liquidscheme":    trueStr,
// 			"overnightscheme": falseStr,
// 			"extendedhours":   falseStr,
// 		},
// 	}
// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Properties: map[string]string{"amfiordercutoff": "1330", "Novaordercutoff": "1300",
// 				"fundscutoff": "1230"},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)
// 	want := ActionSet{
// 		Properties: map[string]string{"amfiordercutoff": "1330", "Novaordercutoff": "1300",
// 			"fundscutoff": "1230"},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"sip liquid order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})

// }

// func testSwitchPhysicalOrder(tests *[]doMatchTest) {

// 	entity := Entity{
// 		Realm: "1",
// 		App:   "Test9",
// 		Slice: "9",
// 		Class: orderClass,

// 		Attrs: map[string]string{
// 			"ordertype":       "switch",
// 			"mode":            "physical",
// 			"liquidscheme":    falseStr,
// 			"overnightscheme": trueStr,
// 			"extendedhours":   trueStr,
// 		},
// 	}

// 	rule2 := Rule_t{
// 		RulePatterns: []RulePatternBlock_t{},
// 		RuleActions: RuleActionBlock_t{
// 			Task:       []string{},
// 			Properties: map[string]string{},
// 		},
// 	}

// 	rs := Ruleset_t{
// 		Id:      1,
// 		Class:   orderClass,
// 		SetName: mainRS,
// 		Rules:   []Rule_t{rule2},
// 		NCalled: 0,
// 	}
// 	ruleSetsTests = append(ruleSetsTests, &rs)

// 	want := ActionSet{
// 		Tasks:      []string{},
// 		Properties: map[string]string{},
// 	}
// 	*tests = append(*tests, doMatchTest{
// 		"switch physical order",
// 		entity,
// 		&rs,
// 		ActionSet{
// 			Tasks:      []string{},
// 			Properties: make(map[string]string),
// 		},
// 		want,
// 	})
// }
