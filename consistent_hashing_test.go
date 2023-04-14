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

	if actualValue != "key-example" {
		t.Errorf("value should be: %s but it is: %s", "key-example", actualValue)
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

	if actualValue != "server-example" {
		t.Errorf("value should be: %s but it is: %s", "server-example", actualValue)
	}
}
