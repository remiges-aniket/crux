// /* This file contains TestMatchPattern() and one helper function */

package crux

// import (
// 	"testing"
// 	"time"
// )

// func TestMatchPattern(t *testing.T) {
// 	var testNames []string
// 	var entities = sampleEntity
// 	var rulePatterns []([]RulePatternBlock_t)
// 	var resultsExpected []any

// 	actionSet := ActionSet{Tasks: []string{"dodiscount", "yearendsale"}}

// 	// Test: many terms, everything matches
// 	testNames = append(testNames, "everything matches")

// 	//receivedTime, _ := time.Parse(timeLayout, "2018-05-15T12:00:00Z")
// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"cat", opEQ, "textbook"},
// 		{"fullname", opEQ, "Advanced Physics"},
// 		{"ageinstock", opLE, 7},
// 		{"mrp", opLT, 51.2},
// 		//{"received", opGT, receivedTime},
// 		{"bulkorder", opNE, false},
// 		{"dodiscount", opEQ, true},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: many terms, mrp doesn't match
// 	testNames = append(testNames, "mrp doesn't match")

// 	receivedTime, _ := time.Parse(timeLayout, "2018-05-15T12:00:00Z")
// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"cat", opEQ, "textbook"},
// 		{"fullname", opEQ, "Advanced Physics"},
// 		{"ageinstock", opLE, 7},
// 		{"mrp", opGE, 51.2},
// 		{"received", opGT, receivedTime},
// 		{"bulkorder", opNE, false},
// 		{"dodiscount", opEQ, true},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	// Test: bool "ne"
// 	testNames = append(testNames, "bool ne")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"bulkorder", opNE, true},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	// Test: enum "ne"
// 	testNames = append(testNames, "enum ne")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"cat", opNE, "refbook"},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: float "eq"
// 	testNames = append(testNames, "float eq")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"mrp", opEQ, 50.8},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: float "ge"
// 	testNames = append(testNames, "float ge")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"mrp", opGE, 50.8},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: timestamp "lt"
// 	testNames = append(testNames, "timestamp lt")

// 	receivedTime, _ = time.Parse(timeLayout, "2018-06-10T15:04:05Z")
// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"received", opLT, receivedTime},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: timestamp "le"
// 	testNames = append(testNames, "timestamp le")

// 	receivedTime, _ = time.Parse(timeLayout, "2018-05-01T15:04:05Z")
// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"received", opLE, receivedTime},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	// Test: string "lt"
// 	testNames = append(testNames, "string lt")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"fullname", opLT, "Advanced Science"},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: string "gt"
// 	testNames = append(testNames, "string gt")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"fullname", opGT, "Accelerated Physics"},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: task wanted in pattern, found in action set
// 	testNames = append(testNames, "Tasks found in action set")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"dodiscount", opEQ, true},
// 		{"yearendsale", opNE, false},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: task wanted in pattern, but not found in action set
// 	testNames = append(testNames, "task not in action set")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"dodiscount", opEQ, true},
// 		{"summersale", opEQ, true},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	// Test: task not wanted in pattern, and not found in action set
// 	testNames = append(testNames, "task 'eq false' in pattern, and not in action set")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"summersale", opEQ, false},
// 	})
// 	resultsExpected = append(resultsExpected, true)

// 	// Test: edge case - no rule pattern
// 	testNames = append(testNames, "no rule pattern")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{})
// 	resultsExpected = append(resultsExpected, true)

// 	// Run the tests
// 	t.Log("==Running", len(rulePatterns), "matchPattern tests==")
// 	ruleSchemas, _ := retriveRuleSchemasAndRuleSetsFromCache(entities.Realm, entities.App, entities.Class, entities.Slice)
// 	for i, rulePattern := range rulePatterns {

// 		res, err := matchPattern(entities, rulePattern, actionSet, ruleSchemas)

// 		if resultsExpected[i] == nil && err == nil {
// 			t.Errorf("Expected but did not get error")
// 			continue
// 		} else if resultsExpected[i] == nil && err != nil {
// 			continue
// 		} else if res != resultsExpected[i] {
// 			t.Errorf(" FAIL output=%v, expected=%v testNames[i] =%v", res, resultsExpected[i], testNames[i])

// 		}
// 	}
// 	testMatchPatternMore(t)
// }
// func testMatchPatternMore(t *testing.T) {
// 	var testNames []string

// 	var rulePatterns []([]RulePatternBlock_t)
// 	var resultsExpected []any
// 	actionSet := ActionSet{Tasks: []string{}}
// 	// Test: error converting value

// 	var sampleEntityrecheck = Entity{
// 		Realm: "1",
// 		App:   "Test6",
// 		Slice: "6",
// 		Class: inventoryItemClass,
// 		Attrs: map[string]string{
// 			"ageinstock": "abc",
// 		},
// 	}
// 	var entities = sampleEntityrecheck
// 	testNames = append(testNames, "deliberate error converting value")
// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"ageinstock", opGT, 5},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	// Test: error - not an ordered type
// 	testNames = append(testNames, "deliberate error: not an ordered type")

// 	rulePatterns = append(rulePatterns, []RulePatternBlock_t{
// 		{"bulkorder", opGT, true},
// 	})
// 	resultsExpected = append(resultsExpected, false)

// 	ruleSchema, _ := retriveRuleSchemasAndRuleSetsFromCache(sampleEntityrecheck.Realm, sampleEntityrecheck.App, sampleEntityrecheck.Class, sampleEntityrecheck.Slice)
// 	for i, rulePattern := range rulePatterns {

// 		res, err := matchPattern(entities, rulePattern, actionSet, ruleSchema)

// 		if resultsExpected[i] == false && err == nil {
// 			t.Errorf("Expected but did not get error")
// 			continue
// 		} else if resultsExpected[i] == nil && err != nil {
// 			continue
// 		} else if res != resultsExpected[i] {
// 			t.Errorf(" FAIL output=%v, expected=%v testNames[i] =%v", res, resultsExpected[i], testNames[i])

// 		}
// 	}
// }

// func setupInventoryItemSchema() {
// 	rs := &Schema_t{
// 		Class: transactionClass,
// 		PatternSchema: []PatternSchema_t{
// 			{Attr: "cat", ValType: typeEnum},
// 			{Attr: "fullname", ValType: typeStr},
// 			{Attr: "ageinstock", ValType: typeInt},
// 			{Attr: "mrp", ValType: typeFloat},
// 			{Attr: "received", ValType: typeTS},
// 			{Attr: "bulkorder", ValType: typeBool},
// 		},
// 		ActionSchema: ActionSchema_t{},
// 	}
// 	ruleSchemasTest = append(ruleSchemasTest, rs)
// }
