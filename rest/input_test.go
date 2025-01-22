package rest

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestReadPVOK(t *testing.T) {
	f, err := os.Open("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	i := Input{&SynkObject{}}
	err = json.Unmarshal(data, i.SynkObject)
	if err != nil {
		t.Fatal(err)
	}
	pv, _ := i.PV(0)
	exp := "2025-01-22 17:55:06"
	if pv["time"] != exp {
		t.Errorf("Expected %s, got %v", exp, pv["time"])
	}
	pv, _ = i.PV(1)
	if pv["time"] != exp {
		t.Errorf("Expected %s, got %v", exp, pv["time"])
	}
}

func TestReadPVOutOfRange(t *testing.T) {
	f, err := os.Open("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	i := Input{&SynkObject{}}
	err = json.Unmarshal(data, i.SynkObject)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := i.PV(2)
	if ok {
		t.Error("expected value to be out of range")
	}
}
