package main

import "testing"

func TestConsistentHashing_PlotKey_ShouldPlotKeyOnTheRingAsPerTheHashValue(t *testing.T) {
	testConsistentHashing := NewConsistentHashing(1000)

	expectedHashKey := 860
	actualHashKey := testConsistentHashing.PlotKey("key-example")

	if actualHashKey != expectedHashKey {
		t.Errorf("hash key should be: %d but it is: %d", expectedHashKey, actualHashKey)
	}

	actualValue := testConsistentHashing.getHashValue(expectedHashKey)

	if actualObj, ok := actualValue.(Key); ok {
		if actualObj.name != "key-example" {
			t.Errorf("value should be: %s but it is: %s", "key-example", actualValue)
		}
	} else {
		t.Errorf("value is not of Key type: %v", actualValue)
	}
}

func TestConsistentHashing_PlotServer_ShouldPlotServerOnTheRingAsPerTheHashValue(t *testing.T) {
	testConsistentHashing := NewConsistentHashing(1000)

	expectedHashKey := 74
	actualHashKey := testConsistentHashing.PlotServer("server-example")

	if actualHashKey != expectedHashKey {
		t.Errorf("hash key should be: %d but it is: %d", expectedHashKey, actualHashKey)
	}

	actualValue := testConsistentHashing.getHashValue(expectedHashKey)

	if actualObj, ok := actualValue.(Server); ok {
		if actualObj.name != "server-example" {
			t.Errorf("value should be: %s but it is: %s", "server-example", actualValue)
		}
	} else {
		t.Errorf("value is not of server type: %v", actualValue)
	}
}
