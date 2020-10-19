package strcycle

import (
	"log"
	"testing"
)

func assertAll(results []bool) (hasError bool) {
	for _, val := range results {
		if val == false {
			return true
		}
	}
	return false
}

func TestCycle(t *testing.T) {
	c, err := Cycle(
		[]interface{}{"a", "b", "c", "d", "e"},
		true,
	)
	if err != nil {
		t.Error(err)
	}
	results := []string{}
	for i := 0; i < 10; i++ {
		results = append(results, (<-c.channel).(string))
	}

	if assertAll(
		[]bool{
			results[0] == "a",
			results[1] == "b",
			results[2] == "c",
			results[3] == "d",
			results[4] == "e",
			results[5] == "**FINISHED**",
			results[6] == "a",
		},
	) {
		t.Fail()
	}
}

func TestFromStringWithModifiersUpperLower(t *testing.T) {
	c, err := FromStringWithModifiers("hey", []StringModifier{UPPER, LOWER}, true)
	if err != nil {
		t.Error(err)
	}
	results := []string{}
	for i := 0; i < 18; i++ {
		results = append(results, (<-c.channel).(string))
	}

	if assertAll(
		[]bool{
			results[6] == "hey",
			results[7] == "HEY",
			results[8] == "**FINISHED**",
		},
	) {
		t.Fail()
	}
}

func TestChain(t *testing.T) {
	c, err := ChainStringArrayWithModifiers(
		[]string{"hello", "world", "yes"},
		[]StringModifier{UPPER, LOWER},
		true,
	)
	if err != nil {
		t.Error(err)
	}
	results := []string{}
	for i := 0; i < 200; i++ {
		v := (<-c.channel).(string)
		log.Println(v)
		results = append(results, v)
	}
	if assertAll(
		[]bool{
			results[72] == "**FINISHED**",
			results[145] == "**FINISHED**",
		},
	) {
		t.Fail()
	}
}

func TestFormatString(t *testing.T) {
	base := "www %s aaa bbb ccc %s %s"
	regions, err := ChainStringArrayWithModifiers(
		[]string{"AU", "GB", "US", "MX"},
		[]StringModifier{UPPER},
		true,
	)
	if err != nil {
		t.Error(err)
	}
	names, err := FromStringArray([]string{"Daniel", "Chloe", "Max", "Adam", "Arnold"}, true)
	if err != nil {
		t.Error(err)
	}
	apple, err := FromStringWithModifiers("apple", []StringModifier{SingleEncodeLower, LOWER}, true)
	if err != nil {
		t.Error(err)
	}

	c, err := FormatStringWithCycles(base, regions, names, apple)
	if err != nil {
		t.Error(err)
	}
	// c := FormatStringWithCycles(base, apple)
	// results := []string{}
	for i := 0; i < 200; i++ {
		v := (<-c.channel).(string)
		log.Println(v)
		// results = append(results, v)
	}
	// log.Println(results)
	t.Fail()
}
