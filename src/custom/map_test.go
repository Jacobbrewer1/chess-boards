package custom

import (
	"testing"
)

func TestMap_Set(t *testing.T) {
	// Test 1
	var key, value = "testOne", "testsOneValue"
	var testMap Map[string, string]

	if testMap != nil {
		t.Error("map not equal to nil")
		return
	}

	testMap.Set(key, value)

	if len(testMap) != 1 {
		t.Error("map length != 1")
		return
	}

	testMap.Set(value, key)

	if len(testMap) != 2 {
		t.Error("map length != 2")
		return
	}

	// Test 2
	var keyTwo = 11
	var testMapTwo Map[int, bool]

	if testMapTwo != nil {
		t.Error("mapTwo not equal to nil")
		return
	}

	testMapTwo.Set(keyTwo, true)

	if len(testMapTwo) != 1 {
		t.Error("mapTwo length != 1")
		return
	}

	// Test 3
	// Test 2
	var keyThree = 11
	var testMapThree Map[int, bool]

	if testMapThree != nil {
		t.Error("mapThree not equal to nil")
		return
	}

	testMapThree.Set(keyThree, true)

	if len(testMapThree) != 1 {
		t.Error("mapTwo length != 1")
		return
	}

	got := testMapThree.Get(keyThree)
	if got == nil {
		t.Error("keyThree not found in map")
		return
	} else if *got != true {
		t.Error("got value not equal to set")
		return
	}

	testMapThree.Set(keyThree, false)
	got = testMapThree.Get(keyThree)
	if got == nil {
		t.Error("keyThree not found in map after reset")
		return
	} else if *got != false {
		t.Error("got value not equal to set after reset")
		return
	}
}

func TestMap_Get(t *testing.T) {
	var testMap Map[string, string]

	testMap.Set("one", "valueOne")
	testMap.Set("two", "valueTwo")
	testMap.Set("three", "valueThree")
	testMap.Set("four", "valueFour")
	testMap.Set("five", "valueFive")

	if got := testMap.Get("six"); got != nil {
		t.Error("'six' exists without being set")
		return
	}

	if got := testMap.Get("three"); got == nil {
		t.Error("three value not found after being set")
		return
	} else if *got != "valueThree" {
		t.Error("three value not equal to value set")
		return
	}
}
