package rest

import "testing"

func TestUpdateOK(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foo", 5)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if object["foo"] != 5 {
		t.Errorf("Expected foo=5, but got %v", object["foo"])
	}
	err = object.Update("bar", "hello")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if object["bar"] != "hello" {
		t.Errorf("Expected bar=baz, but got %v", object["bar"])
	}
}

func TestUpdateInvalidKey(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foobar", 5)
	if err == nil {
		t.Error("updated non-existent key")
	}
}

func TestUpdateWrongTypes(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foo", "string")
	if err == nil {
		t.Errorf("updated foo value with wrong type")
	}
	err = object.Update("bar", 12345678)
	if err == nil {
		t.Errorf("updated bar value with wrong type")
	}
}
