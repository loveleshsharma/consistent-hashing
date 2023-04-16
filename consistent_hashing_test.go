package main

import "testing"

func TestConsistentHashing_PlotKey_ShouldPlotKeyOnTheRingAsPerTheHashValue(t *testing.T) {
	testConsistentHashing := NewConsistentHashing(1000)

	expectedHashKey := 718
	actualHashKey := testConsistentHashing.PlotKey(NewKey("key-example"))

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

	expectedHashKey := 29
	actualHashKey := testConsistentHashing.PlotServer(NewServer("server-example"))

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

func TestConsistentHashing_GetServer_ShouldReturnServerForTheKey(t *testing.T) {
	testConsistentHashing := NewConsistentHashing(50)

	testConsistentHashing.PlotServer(NewServer("s1"))
	testConsistentHashing.PlotKey(NewKey("k1"))
	testConsistentHashing.PlotKey(NewKey("k2"))

	expectedServer := "s1"
	actualServer := testConsistentHashing.GetServer(NewKey("k1"))

	if actualServer.name != expectedServer {
		t.Errorf("server should be: %s but it is %s", expectedServer, actualServer.name)
	}

	actualServer = testConsistentHashing.GetServer(NewKey("k2"))

	if actualServer.name != expectedServer {
		t.Errorf("server should be: %s but it is %s", expectedServer, actualServer.name)
	}
}
