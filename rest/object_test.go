package rest

import "testing"

func TestExtractOk(t *testing.T) {
	o := SynkObject{}
	o["foo"] = 3
	o["bar"] = 7
	o["hello"] = "world"
	keys := []string{"foo", "hello"}
	subset, err := o.ExtractKeys(keys)
	if err != nil {
		t.Errorf("test failed with reason %v", err)
	}
	if len(*subset) != 2 {
		t.Errorf("Expected 2 attributes, but got %d", len(*subset))
	}
	if (*subset)["foo"] != 3 {
		t.Errorf("Expected 3, but got %v", (*subset)["foo"])
	}
	if (*subset)["hello"] != "world" {
		t.Errorf("Expected hello, but got %v", (*subset)["hello"])
	}
}

func TestNoSuchKey(t *testing.T) {
	o := SynkObject{"foo": 3}
	subset, err := o.ExtractKeys([]string{"bar"})
	if err == nil {
		t.Error("expected to get an error")
	}
	if subset != nil {
		t.Error("no SynkObject should have been returned")
	}
}
